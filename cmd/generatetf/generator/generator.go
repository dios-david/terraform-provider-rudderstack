package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zclconf/go-cty/cty"

	"github.com/rudderlabs/rudder-api-go/client"
	"github.com/rudderlabs/terraform-provider-rudderstack/rudderstack/configs"
	_ "github.com/rudderlabs/terraform-provider-rudderstack/rudderstack/integrations"
)

var logger = log.New(os.Stderr, "", 0)

type sourceEntry struct {
	terraformType string
	source        client.Source
}

type destinationEntry struct {
	terraformType string
	destination   client.Destination
}

// GenerateImportScript generates terraform import commands that import resources from terraform script as generated by GenerateTerraform.
func GenerateImportScript(sources []client.Source, destinations []client.Destination, connections []client.Connection) ([]byte, error) {
	lines := []string{}

	sourceConfigs := configs.Sources.Entries()
	destinationConfigs := configs.Destinations.Entries()

	foundSources := map[string]bool{}
	foundDestinations := map[string]bool{}

	for _, src := range sources {
		t, cm := configMeta(sourceConfigs, src.Type)
		if cm != nil {
			foundSources[src.ID] = true
			lines = append(lines, fmt.Sprintf(`terraform import "rudderstack_source_%s.%s" "%s"`, t, sourceName(src), src.ID))
		}
	}

	for _, dst := range destinations {
		t, cm := configMeta(destinationConfigs, dst.Type)
		if cm != nil {
			foundDestinations[dst.ID] = true
			lines = append(lines, fmt.Sprintf(`terraform import "rudderstack_destination_%s.%s" "%s"`, t, destinationName(dst), dst.ID))
		}
	}

	for _, cnxn := range connections {
		_, oks := foundSources[cnxn.SourceID]
		_, okd := foundDestinations[cnxn.DestinationID]
		if oks && okd {
			lines = append(lines, fmt.Sprintf(`terraform import "rudderstack_connection.%s" "%s"`, connectionName(cnxn), cnxn.ID))
		}
	}

	script := strings.Join(lines, "\n")
	return []byte(script), nil
}

// Generate HCL script from a given set of API sources, destinations and connections.
func GenerateTerraform(sources []client.Source, destinations []client.Destination, connections []client.Connection) ([]byte, error) {
	f := hclwrite.NewEmptyFile()
	body := f.Body()

	// generate source blocks for supported sources
	sourceConfigs := configs.Sources.Entries()
	generatedSourceEntries := map[string]*sourceEntry{}
	for _, src := range sources {
		terraformType, cm := configMeta(sourceConfigs, src.Type)
		if cm != nil {
			b, err := generateSource(src, terraformType, cm)
			if err != nil {
				return nil, err
			}

			body.AppendBlock(b)
			body.AppendNewline()
			generatedSourceEntries[src.ID] = &sourceEntry{terraformType: terraformType, source: src}
		} else {
			logger.Printf("could not generate resource block for source '%s':  type '%s' is not supported", src.ID, src.Type)
		}
	}

	// generate destination blocks for supported destinations
	destinationConfigs := configs.Destinations.Entries()
	generatedDestinationEntries := map[string]*destinationEntry{}
	for _, dst := range destinations {
		terraformType, cm := configMeta(destinationConfigs, dst.Type)
		if cm != nil {
			b, err := generateDestination(dst, terraformType, cm)
			if err != nil {
				return nil, err
			}

			body.AppendBlock(b)
			body.AppendNewline()
			generatedDestinationEntries[dst.ID] = &destinationEntry{terraformType: terraformType, destination: dst}
		} else {
			logger.Printf("could not generate resource block for destination '%s':  type '%s' is not supported", dst.ID, dst.Type)
		}
	}

	// generate connection blocks for supported sources and destinations
	for _, cnxn := range connections {
		src, oks := generatedSourceEntries[cnxn.SourceID]
		dst, okd := generatedDestinationEntries[cnxn.DestinationID]
		if !oks {
			logger.Printf("could not generate resource block for connection '%s': connected source '%s' is not supported", cnxn.ID, cnxn.SourceID)
		} else if !okd {
			logger.Printf("could not generate resource block for connection '%s': connected destination '%s' is not supported", cnxn.ID, cnxn.DestinationID)
		} else {
			b, err := generateConnection(cnxn, src, dst)
			if err != nil {
				return nil, fmt.Errorf("could not generate resource block for connection '%s': %w", cnxn.ID, err)
			}

			body.AppendBlock(b)
			body.AppendNewline()
		}
	}

	return f.Bytes(), nil
}

// generateSource generates a source resource block from an API source object and a terraform source type and ConfigMeta.
func generateSource(source client.Source, terraformType string, cm *configs.ConfigMeta) (*hclwrite.Block, error) {
	resourceType := sourceType(terraformType)
	resourceName := sourceName(source)
	block := hclwrite.NewBlock("resource", []string{resourceType, resourceName})

	body := block.Body()
	body.SetAttributeValue("name", cty.StringVal(source.Name))

	if !cm.SkipConfig {
		configBlock, err := generateConfigBlock(source.Config, cm)
		if err != nil {
			return nil, fmt.Errorf("could not generate config block for source '%s': %w", source.ID, err)
		}
		body.AppendBlock(configBlock)
	}

	return block, nil
}

func sourceType(terraformType string) string {
	return fmt.Sprintf("rudderstack_source_%s", terraformType)
}

func sourceName(source client.Source) string {
	return fmt.Sprintf("src_%s", source.ID)
}

// generateDestination genertes a destination resource block from an API destinaton object and a terraform destination type and ConfigMeta.
func generateDestination(destination client.Destination, terraformType string, cm *configs.ConfigMeta) (*hclwrite.Block, error) {
	resourceType := destinationType(terraformType)
	resourceName := destinationName(destination)
	block := hclwrite.NewBlock("resource", []string{resourceType, resourceName})

	body := block.Body()
	body.SetAttributeValue("name", cty.StringVal(destination.Name))

	if !cm.SkipConfig {
		configBlock, err := generateConfigBlock(destination.Config, cm)
		if err != nil {
			return nil, fmt.Errorf("could not generate config block for destination '%s': %w", destination.ID, err)
		}
		body.AppendBlock(configBlock)
	}

	return block, nil
}

func destinationType(terraformType string) string {
	return fmt.Sprintf("rudderstack_destination_%s", terraformType)
}

func destinationName(destination client.Destination) string {
	return fmt.Sprintf("dst_%s", destination.ID)
}

// configMeta finds a ConfigMeta of a specific api type. Returns the terraform type and the ConfigMeta if found.
// if not, it returns an empty string and nil.
func configMeta(entries map[string]configs.ConfigMeta, apiType string) (string, *configs.ConfigMeta) {
	for r, e := range entries {
		if e.APIType == apiType {
			return r, &e
		}
	}

	return "", nil
}

// generateConfigBlock generate a source or destination config block from an API config JSON
// and a corresponding ConfigMeta.
func generateConfigBlock(config json.RawMessage, cm *configs.ConfigMeta) (*hclwrite.Block, error) {
	// get state representation of config as map[string]interface{}
	state, err := cm.APIToState(string(config))
	if err != nil {
		return nil, fmt.Errorf("could not convert API config to terraform state: %w", err)
	}

	var stateMap map[string]interface{}
	if err := json.Unmarshal([]byte(state), &stateMap); err != nil {
		return nil, err
	}

	block, err := generateBlock("config", stateMap, cm.ConfigSchema)
	if err != nil {
		return nil, fmt.Errorf("could not generate config block: %w", err)
	}

	return block, nil
}

// generateBlock generate blocks with specified name and data which conform to the provided config schema.
// data should be unmarshaled from a state JSON object.
func generateBlock(name string, data map[string]interface{}, configSchema map[string]*schema.Schema) (*hclwrite.Block, error) {
	block := hclwrite.NewBlock(name, []string{})
	body := block.Body()

	// go does not garantee the order of range in maps, we sort the keys so that the output is predictable
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := data[k]
		if sch, ok := configSchema[k]; ok {
			// state objects are actually lists with a single object item
			if r, ok := sch.Elem.(*schema.Resource); ok && sch.Type == schema.TypeList && sch.ConfigMode != schema.SchemaConfigModeAttr {
				if l, ok := v.([]interface{}); ok && len(l) > 0 {
					if value, ok := l[0].(map[string]interface{}); ok {
						kBlock, err := generateBlock(k, value, r.Schema)
						if err != nil {
							return nil, err
						}
						body.AppendBlock(kBlock)
					}
				}
			} else {
				body.SetAttributeValue(k, ctyValue(v))
			}
		}
	}

	return block, nil
}

// ctyValue converts any value to a cty.Value which can be passed as values to HCL attributes
func ctyValue(x interface{}) cty.Value {
	switch v := x.(type) {
	case string:
		return cty.StringVal(v)

	case bool:
		return cty.BoolVal(v)

	case int:
	case int16:
	case int32:
	case int64:
		return cty.NumberIntVal(v)

	case float32:
	case float64:
		return cty.NumberFloatVal(v)

	case []interface{}:
		var values []cty.Value
		for _, i := range v {
			values = append(values, ctyValue(i))
		}
		return cty.ListVal(values)

	case map[string]interface{}:
		values := map[string]cty.Value{}
		for k, i := range v {
			values[k] = ctyValue(i)
		}
		return cty.ObjectVal(values)
	}

	return cty.EmptyObjectVal
}

func generateConnection(connection client.Connection, source *sourceEntry, destination *destinationEntry) (*hclwrite.Block, error) {
	resourceType := "rudderstack_connection"
	resourceName := connectionName(connection)
	block := hclwrite.NewBlock("resource", []string{resourceType, resourceName})

	body := block.Body()

	body.SetAttributeRaw("source_id", hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(fmt.Sprintf("%s.%s.id", sourceType(source.terraformType), sourceName(source.source))),
		},
	})

	body.SetAttributeRaw("destination_id", hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(fmt.Sprintf("%s.%s.id", destinationType(destination.terraformType), destinationName(destination.destination))),
		},
	})

	return block, nil
}

func connectionName(connection client.Connection) string {
	return fmt.Sprintf("cnxn_%s", connection.ID)
}

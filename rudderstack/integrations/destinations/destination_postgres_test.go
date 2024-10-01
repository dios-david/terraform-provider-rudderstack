package destinations_test

import (
	"testing"

	cmt "github.com/rudderlabs/terraform-provider-rudderstack/internal/testutil/cm"
	c "github.com/rudderlabs/terraform-provider-rudderstack/rudderstack/configs"
)

func TestDestinationResourcePostgres(t *testing.T) {
	cmt.AssertDestination(t, "postgres", []c.TestConfig{
		{
			TerraformCreate: `
				host = "test-host"
				database = "test-database"
				user = "test-user"
				password = "test-password"
				port = "test-port"
				ssl_mode = "disable"
				sync_frequency = "30"
				use_rudder_storage = true
			`,
			APICreate: `{
				"host": "test-host",
				"database": "test-database",
				"user": "test-user",
				"password": "test-password",
				"port": "test-port",
				"sslMode": "disable",
				"syncFrequency": "30",
				"useRudderStorage": true
			}`,
			TerraformUpdate: `
				host = "test-host"
				database = "test-database"
				user = "test-user"
				password = "test-password"
				port = "test-port"
				ssl_mode = "verify-ca"
				sync_frequency = "60"
				use_rudder_storage = true
				consent_management {
					web = [
						{
							provider = "oneTrust"
							consents = ["one_web", "two_web", "three_web"]
							resolution_strategy = ""
						},
						{
							provider = "ketch"
							consents = ["one_web", "two_web", "three_web"]
							resolution_strategy = ""
						},
						{
							provider = "custom"
							resolution_strategy = "and"
							consents = ["one_web", "two_web", "three_web"]
						}
					]
					android = [{
						provider = "ketch"
						consents = ["one_android", "two_android", "three_android"]
						resolution_strategy = ""
					}]
					ios = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_ios", "two_ios", "three_ios"]
					}]
					unity = [{
						provider = "custom"
						resolution_strategy = "or"
						consents = ["one_unity", "two_unity", "three_unity"]
					}]
					reactnative = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_reactnative", "two_reactnative", "three_reactnative"]
					}]
					flutter = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_flutter", "two_flutter", "three_flutter"]
					}]
					cordova = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_cordova", "two_cordova", "three_cordova"]
					}]
					amp = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_amp", "two_amp", "three_amp"]
					}]
					cloud = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_cloud", "two_cloud", "three_cloud"]
					}]
					cloud_source = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_cloud_source", "two_cloud_source", "three_cloud_source"]
					}]
					shopify = [{
						provider = "custom"
						resolution_strategy = "and"
						consents = ["one_shopify", "two_shopify", "three_shopify"]
					}]
				}
			`,
			APIUpdate: `{
				"host": "test-host",
				"database": "test-database",
				"user": "test-user",
				"password": "test-password",
				"port": "test-port",
				"sslMode": "verify-ca",
				"syncFrequency": "60",
				"useRudderStorage": true,
				"consentManagement": {
					"web": [
						{
							"provider": "oneTrust",
							"resolutionStrategy": "",
							"consents": [
								{
									"consent": "one_web"
								},
								{
									"consent": "two_web"
								},
								{
									"consent": "three_web"
								}
							]
						},
						{
							"provider": "ketch",
							"resolutionStrategy": "",
							"consents": [
								{
									"consent": "one_web"
								},
								{
									"consent": "two_web"
								},
								{
									"consent": "three_web"
								}
							]
						},
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_web"
								},
								{
									"consent": "two_web"
								},
								{
									"consent": "three_web"
								}
							]
						}
					],
					"android": [
						{
							"provider": "ketch",
							"resolutionStrategy": "",
							"consents": [
								{
									"consent": "one_android"
								},
								{
									"consent": "two_android"
								},
								{
									"consent": "three_android"
								}
							]
						}
					],
					"ios": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_ios"
								},
								{
									"consent": "two_ios"
								},
								{
									"consent": "three_ios"
								}
							]
						}
					],
					"unity": [
						{
							"provider": "custom",
							"resolutionStrategy": "or",
							"consents": [
								{
									"consent": "one_unity"
								},
								{
									"consent": "two_unity"
								},
								{
									"consent": "three_unity"
								}
							]
						}
					],
					"reactnative": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_reactnative"
								},
								{
									"consent": "two_reactnative"
								},
								{
									"consent": "three_reactnative"
								}
							]
						}
					],
					"flutter": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_flutter"
								},
								{
									"consent": "two_flutter"
								},
								{
									"consent": "three_flutter"
								}
							]
						}
					],
					"cordova": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_cordova"
								},
								{
									"consent": "two_cordova"
								},
								{
									"consent": "three_cordova"
								}
							]
						}
					],
					"amp": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_amp"
								},
								{
									"consent": "two_amp"
								},
								{
									"consent": "three_amp"
								}
							]
						}
					],
					"cloud": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_cloud"
								},
								{
									"consent": "two_cloud"
								},
								{
									"consent": "three_cloud"
								}
							]
						}
					],
					"cloudSource": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_cloud_source"
								},
								{
									"consent": "two_cloud_source"
								},
								{
									"consent": "three_cloud_source"
								}
							]
						}
					],
					"shopify": [
						{
							"provider": "custom",
							"resolutionStrategy": "and",
							"consents": [
								{
									"consent": "one_shopify"
								},
								{
									"consent": "two_shopify"
								},
								{
									"consent": "three_shopify"
								}
							]
						}
					]
				}
			}`,
		},
	})
}

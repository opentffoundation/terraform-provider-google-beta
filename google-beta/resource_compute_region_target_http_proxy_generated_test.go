// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputeRegionTargetHttpProxy_regionTargetHttpProxyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionTargetHttpProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyBasicExample(context),
			},
		},
	})
}

func testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_target_http_proxy" "default" {
  provider = google-beta

  region  = "us-central1"
  name    = "tf-test-test-proxy%{random_suffix}"
  url_map = google_compute_region_url_map.default.self_link
}

resource "google_compute_region_url_map" "default" {
  provider = google-beta

  region          = "us-central1"
  name            = "tf-test-url-map%{random_suffix}"
  default_service = google_compute_region_backend_service.default.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.default.self_link

    path_rule {
      paths   = ["/*"]
      service = google_compute_region_backend_service.default.self_link
    }
  }
}

resource "google_compute_region_backend_service" "default" {
  provider = google-beta

  region      = "us-central1"
  name        = "tf-test-backend-service%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
}

resource "google_compute_region_health_check" "default" {
  provider = google-beta

  region = "us-central1"
  name   = "tf-test-http-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionTargetHttpProxy_regionTargetHttpProxyHttpsRedirectExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionTargetHttpProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyHttpsRedirectExample(context),
			},
		},
	})
}

func testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyHttpsRedirectExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_target_http_proxy" "default" {
  provider = google-beta

  region  = "us-central1"
  name    = "tf-test-test-https-redirect-proxy%{random_suffix}"
  url_map = google_compute_region_url_map.default.self_link
}

resource "google_compute_region_url_map" "default" {
  provider = google-beta

  region          = "us-central1"
  name            = "tf-test-url-map%{random_suffix}"
  default_url_redirect {
    https_redirect = true
  }
}
`, context)
}

func testAccCheckComputeRegionTargetHttpProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_target_http_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeRegionTargetHttpProxy still exists at %s", url)
			}
		}

		return nil
	}
}

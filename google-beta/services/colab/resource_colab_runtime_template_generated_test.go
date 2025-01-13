// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package colab_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccColabRuntimeTemplate_colabRuntimeTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckColabRuntimeTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccColabRuntimeTemplate_colabRuntimeTemplateBasicExample(context),
			},
			{
				ResourceName:            "google_colab_runtime_template.runtime-template",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccColabRuntimeTemplate_colabRuntimeTemplateBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "runtime-template" {
  name = "tf-test-colab-runtime-template%{random_suffix}"
  display_name = "Runtime template basic"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}
`, context)
}

func TestAccColabRuntimeTemplate_colabRuntimeTemplateNoNameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckColabRuntimeTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccColabRuntimeTemplate_colabRuntimeTemplateNoNameExample(context),
			},
		},
	})
}

func testAccColabRuntimeTemplate_colabRuntimeTemplateNoNameExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "runtime-template" {
  display_name = "Runtime template no name"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}
`, context)
}

func TestAccColabRuntimeTemplate_colabRuntimeTemplateFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"key_name":      acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckColabRuntimeTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccColabRuntimeTemplate_colabRuntimeTemplateFullExample(context),
			},
			{
				ResourceName:            "google_colab_runtime_template.runtime-template",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccColabRuntimeTemplate_colabRuntimeTemplateFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "my_network" {
  name = "tf-test-colab-test-default%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "my_subnetwork" {
  name   = "tf-test-colab-test-default%{random_suffix}"
  network = google_compute_network.my_network.id
  region = "us-central1"
  ip_cidr_range = "10.0.1.0/24"
}

resource "google_colab_runtime_template" "runtime-template" {
  name        = "tf-test-colab-runtime-template%{random_suffix}"
  display_name = "Runtime template full"
  location    = "us-central1"
  description = "Full runtime template"
  machine_spec {
    machine_type     = "n1-standard-2"
    accelerator_type = "NVIDIA_TESLA_T4"
    accelerator_count = "1"
  }

  data_persistent_disk_spec {
    disk_type    = "pd-standard"
    disk_size_gb = 200
  }

  network_spec {
    enable_internet_access = true
    network = google_compute_network.my_network.id
    subnetwork = google_compute_subnetwork.my_subnetwork.id
  }

  labels = {
    k = "val"
  }

  idle_shutdown_config {
    idle_timeout = "3600s"
  }

  euc_config {
    euc_disabled = true
  }

  shielded_vm_config {
    enable_secure_boot = true
  }

  network_tags = ["abc", "def"]

  encryption_spec {
    kms_key_name = "%{key_name}"
  }
}
`, context)
}

func testAccCheckColabRuntimeTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_colab_runtime_template" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ColabBasePath}}projects/{{project}}/locations/{{location}}/notebookRuntimeTemplates/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ColabRuntimeTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
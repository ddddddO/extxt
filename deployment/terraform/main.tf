provider "google" {
  credentials = "${file("~/.config/gcloud/legacy_credentials/lbfdeatq@gmail.com/adc.json")}"

  project = "extxt-300211"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-c"
}

provider "google-beta" {
  credentials = "${file("~/.config/gcloud/legacy_credentials/lbfdeatq@gmail.com/adc.json")}"

  project = "extxt-300211"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-c"
}

## NOTE: extxtに変更があった場合は、make buildappでイメージを更新&GCRへpushする。で、cloud runをdestroy -> applyする
resource "google_cloud_run_service" "extxt" {
  provider = google-beta
  name     = "extxt-app"
  location = "asia-northeast1"

  template {
    spec {
      containers {
        image = "gcr.io/extxt-300211/extxt"
        env {
          name  = "BASIC_AUTH_NAMES"
          value = "${var.BASIC_AUTH_NAMES}"
        }
        env {
          name  = "BASIC_AUTH_PASSWORDS"
          value = "${var.BASIC_AUTH_PASSWORDS}"
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "1000"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  autogenerate_revision_name = true
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  provider = google-beta
  location = google_cloud_run_service.extxt.location
  project  = google_cloud_run_service.extxt.project
  service  = google_cloud_run_service.extxt.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

output "app_status" {
  value = "${google_cloud_run_service.extxt.status}"
}
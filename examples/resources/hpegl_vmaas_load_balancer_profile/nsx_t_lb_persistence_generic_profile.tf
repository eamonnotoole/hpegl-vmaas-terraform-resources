# (C) Copyright 2022 Hewlett Packard Enterprise Development LP

# PERSISTENCE Profile for GENERIC service
resource "hpegl_vmaas_load_balancer_profile" "tf_PERSISTENCE-GENERIC" {
  lb_id = data.hpegl_vmaas_load_balancer.tf_lb.id  
  name  =  "tf_PERSISTENCE-GENERIC"       
  description  = "PERSISTENCE-GENERIC creating using tf"
  service_type     = "LBGenericPersistenceProfile"
  config {
    generic_profile{
      profile_type = "persistence-profile"
      share_persistence = false
      ha_persistence_mirroring = false
      persistence_entry_timeout = 30
    }
    tags {
        tag = "tag1"
        scope = "scope1"
    }
  }
}
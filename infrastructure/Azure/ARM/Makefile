LOCATION := "Central India"
RGROUP_NAME := "CL01"
TEMPLATE := "azuredeploy.json"

DEPLOYMENT := "devops"
STAGE := "dev"
AKS_CLUSTER := $(DEPLOYMENT)-"aks"-$(STAGE)
DEPLOYMENT := $(DEPLOYMENT)-$(STAGE)
CONFIG_PATH := $(PWD)/kubeconfig

STORAGE_ACCOUNT_NAME := "devops2022dev"

create_resource_group:
	az group create \
		--name  $(RGROUP_NAME) \
		--location $(LOCATION) 

delete_resource_group:
	az group delete \
		-n $(RGROUP_NAME) -y


deploy:
	az deployment group create \
		--name $(DEPLOYMENT) \
		--resource-group $(RGROUP_NAME) \
		--template-file $(TEMPLATE)
		--parameters storageName=$(STORAGE_ACCOUNT_NAME)

delete:
	az deployment group delete \
		-n $(DEPLOYMENT) \
		-g ${RGROUP_NAME}
	
get_aks_config:
	az aks get-credentials \
		--resource-group $(RGROUP_NAME)  \
		--name $(AKS_CLUSTER) \
		--file $(CONFIG_PATH)

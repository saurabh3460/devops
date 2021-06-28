1. This project is to demostrate golang app with PostgresQL.
2. This app will be used to demostrate [autoscaling strategy](https://learnk8s.io/kubernetes-autoscaling-strategies).
3. Will use testing tools to test load on network and cpu, will try to fill the memory and storage
4. Use Helm chart
5. Use ArgoCD for GitOps
6. Write Terraform for azure.
7. Try to automate whole thing by creating cli or script

## Phase One: 

#### 1 step 
Write web app with postgreSQL or redis(optional) and Dockerfile. 

#### 2 Write Helm chart 

#### 3 Write argocd manifest use (ClI to manage)

## Phase Two:

#### Step 1 
Setup logging.

#### Step 2
Setup backup for psql inside cluster.

#### Step 3 (can be contine at any phase)
find tools for network testing or load testing.

## Phase Three:

### Step 1:
Use different autoscaling-strategies.
Can use caos framework.

### Step 2:
Setup monitoring tool

## Phase Four:

### Step 1:
Write Terraform for azure. (k8s)
Find best practice for security and scalling.
# GITHUB ACTIONS WORKFLOWS

We have  2 workflows just to get CI

[ci.yaml](../.github/workflows/ci.yaml): 
- Launch unit test.
- Build an image inside a container.
- Push image to ismaelfm/node-port-scanner
- The tag of the image will be both latest and GITHUB_SHA.

[terratest.yaml](../.github/workflows/terratest.yaml):
- Will launch terratest.
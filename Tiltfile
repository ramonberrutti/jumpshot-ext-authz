version_settings(constraint='>=0.22.2')

docker_build(
    'jumpshot-ext-authz',
    context='.',
    dockerfile='./Dockerfile',
    only=[
        './go.mod', 
        './go.sum', 
        './cmd/',
        './internal/',
        './vendor/',
    ],
    live_update=[

    ]
)

k8s_yaml('k8s/namespace.yaml')
k8s_yaml('k8s/deployment.yaml')
k8s_yaml('k8s/service.yaml')
k8s_yaml('k8s/podmonitor.yaml')

k8s_resource(
    'jumpshot-ext-authz',
    labels=['jumpshot-ext-authz']
)

load('ext://helm_resource', 'helm_resource')
load('ext://helm_resource', 'helm_repo')

helm_repo('postgresql-helm', 'https://charts.bitnami.com/bitnami', labels=["helm"])

helm_resource('postgresql',
            chart='postgresql-helm/postgresql',
            release_name='postgresql',
            resource_deps=['postgresql-helm'],
            labels=["DB"],
            deps=[
            'values.yaml'
            ],
            flags=[
            '--set',  'image.tag=latest',
            '--set',  'nameOverride=hradec-database',
            '--set',  'auth.enablePostgresUser=true',
            '--set',  'auth.postgresPassword=postgres',
            '--set',  'auth.database=hradec',
            ]
)

k8s_resource('postgresql', port_forwards="5434:5432", labels=["DB"])


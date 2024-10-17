load('ext://helm_remote', 'helm_remote')
load('ext://helm_resource', 'helm_resource')
load('ext://helm_resource', 'helm_repo')
load('ext://restart_process', 'docker_build_with_restart')
load('ext://local_output', 'local_output')

load_dynamic('./ci/tilt/postgres.Tiltfile')

k8s_yaml("ci/kube_hradec.yaml")


local_resource(
      'compile hradec',
      'cd hradec && bash ./ci/build.sh',
      deps=[
      './hradec/',
      ],
      ignore=[
      'tilt_modules',
      'Tiltfile',
      'graph/schema.graphqls',
      'hradec/build',
      'dep',
      'ci/docker-compose.yaml',
      'swagger.yaml',
      '**/testdata'
      ],
      labels=["compile"],
  )
  
docker_build_with_restart('hradec',
    '.',
    dockerfile='./hradec/ci/Dockerfile',
    entrypoint='/app/start_server',
    only=[
        './hradec/build',
        './hradec/configurations',
        './hradec/migrations',
    ],
    live_update=[
        sync('./configurations', '/app/configurations'),
        sync('./build', '/app')
    ]
)




name = "hradec"

docker_build_with_restart('hradec-migrations',
           '.',
           dockerfile='hradec/ci/Dockerfile',
           entrypoint='/app/run_migrations',
           only=[
               './%s/build' % name,
               './%s/ci' % name,
               './%s/configurations' % name,
               './%s/certs' % name,
               './%s/migrations' % name
           ],
           live_update=[
               sync('./%s/build' % name , '/app'),
               sync('./%s/configurations' % name , '/app/configurations')
           ],
           build_args={"app": name})

k8s_resource('hradec', port_forwards=["0.0.0.0:8080:8080"], labels=["BE"])


pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    some-label: kaniko-job
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    command:
    - cat
    tty: true
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker
    - name: workspace-volume
      mountPath: /workspace
  volumes:
  - name: docker-config
    secret:
      secretName: dockerhub-secret
      items:
      - key: .dockerconfigjson
        path: config.json
  - name: workspace-volume
    emptyDir: {}
"""
    }
  }

  environment {
    IMAGE = 'bala1511/userapi:latest'
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    sh '''
  echo '📂 Listing contents of /workspace...'
  ls -l /workspace

  echo '🚀 Starting Kaniko Build'
  /kaniko/executor \
    --context=dir:///workspace \
    --dockerfile=/workspace/Dockerfile \
    --destination=docker.io/bala1511/userapi:latest \
    --verbosity=debug
'''

  }
}

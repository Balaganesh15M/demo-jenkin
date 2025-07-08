pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:latest
    tty: true
    command:
    - /kaniko/executor
    args:
    - --dockerfile=Dockerfile
    - --context=dir://workspace
    - --destination=docker.io/bala1115/userapinew:\${BUILD_NUMBER}
    - --destination=docker.io/bala1115/userapinew:latest
    volumeMounts:
      - name: docker-config
        mountPath: /kaniko/.docker
  volumes:
    - name: docker-config
      secret:
        secretName: docker-hub-creds
"""
    }
  }

  stages {
    stage('Build & Push with Kaniko') {
      steps {
        container('kaniko') {
          echo 'Building and pushing image with Kaniko...'
        }
      }
    }
  }
}

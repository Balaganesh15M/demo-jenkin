pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    some-label: kaniko
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:latest
    args:
    - "--dockerfile=Dockerfile"
    - "--context=."
    - "--destination=docker.io/bala1511/go-kaniko-demo:latest"
    - "--verbosity=debug"
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker/
  volumes:
  - name: docker-config
    secret:
      secretName: dockerhub-secret
"""
    }
  }

  stages {
    stage('Checkout Source') {
      steps {
        git branch: 'main', url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
      }
    }

    stage('Build & Push with Kaniko') {
      steps {
        container('kaniko') {
          echo "Kaniko is building and pushing the image..."
          // Nothing to do here because Kaniko runs as entrypoint with args
        }
      }
    }
  }
}

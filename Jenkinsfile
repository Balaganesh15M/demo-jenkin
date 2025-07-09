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
    - "--context=dir://\$(WORKSPACE)"
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

  environment {
    IMAGE = "bala1511/go-kaniko-demo:latest"
  }

  stages {
    stage('Checkout Source') {
      steps {
        git branch: 'main', url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
      }
    }

    stage('Build with Kaniko') {
      steps {
        container('kaniko') {
          echo "Building Docker image with Kaniko..."
          // Kaniko runs with args — no shell command needed
        }
      }
    }
  }
}

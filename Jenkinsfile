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
    image: gcr.io/kaniko-project/executor:v1.9.1
    args:
    - --dockerfile=Dockerfile
    - --context=dir:///workspace
    - --destination=docker.io/${env.IMAGE}
    - --verbosity=debug
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker/
  volumes:
  - name: docker-config
    secret:
      secretName: dockerhub-secret
      items:
      - key: config.json
        path: config.json
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

    stage('Verify Setup') {
      steps {
        container('kaniko') {
          script {
            sh 'ls -la /kaniko/.docker/ || true'
          }
        }
      }
    }

    stage('Build with Kaniko') {
      steps {
        container('kaniko') {
          echo "Kaniko build process starting..."
        }
      }
    }
  }
}

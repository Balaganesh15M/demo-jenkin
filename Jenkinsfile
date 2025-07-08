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

  environment {
    IMAGE = 'bala1115/userapinew'
    TAG = "build-${BUILD_NUMBER}"
  }

  stages {
    stage('Checkout') {
      steps {
        container('kaniko') {
          checkout scm
        }
      }
    }

    stage('Build & Push') {
      steps {
        container('kaniko') {
          sh """
            /kaniko/executor \
              --context `pwd` \
              --dockerfile `pwd`/Dockerfile \
              --destination=$IMAGE:$TAG \
              --destination=$IMAGE:latest
          """
        }
      }
    }
  }
}

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
    command:
    - cat
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
    IMAGE = 'bala1115/userapi'
    TAG = "build-${BUILD_NUMBER}"
  }

  stages {
    stage('Clone Repo') {
      steps {
        container('kaniko') {
          checkout scm
        }
      }
    }

    stage('Build & Push Image') {
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

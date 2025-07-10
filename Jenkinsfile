pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    jenkins: kaniko
spec:
  containers:
    - name: kaniko
      image: gcr.io/kaniko-project/executor:v1.9.1
      args:
        - --dockerfile=/workspace/Dockerfile
        - --context=dir:///workspace
        - --destination=docker.io/bala1511/demo-jenkin:latest
        - --verbosity=debug
        - --skip-tls-verify
      volumeMounts:
        - name: docker-config
          mountPath: /kaniko/.docker
        - name: workspace-volume
          mountPath: /workspace
    - name: jnlp
      image: jenkins/inbound-agent:latest
      volumeMounts:
        - name: workspace-volume
          mountPath: /workspace
  volumes:
    - name: docker-config
      secret:
        secretName: dockerhub-secret
        items:
          - key: config.json
            path: config.json
    - name: workspace-volume
      emptyDir: {}
"""
    }
  }

  stages {
    stage('Clone Repo') {
      steps {
        container('jnlp') {
          dir('/workspace') {
            git url: 'https://github.com/Balaganesh15M/demo-jenkin.git', branch: 'main'
          }
        }
      }
    }
    stage('Build & Push Image') {
      steps {
        // Nothing to do — Kaniko runs automatically using args
        echo 'Kaniko is building and pushing the image...'
      }
    }
  }
}

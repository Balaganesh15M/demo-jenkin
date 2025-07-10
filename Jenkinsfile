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
    command: ["/kaniko/executor"]
    args:
    - "--dockerfile=Dockerfile"
    - "--context=dir:///workspace"
    - "--destination=docker.io/bala1511/demo-jenkin:latest"
    - "--verbosity=debug"
    - "--skip-tls-verify"
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
    stage('Checkout Source') {
      steps {
        git branch: 'main', 
        url: 'https://github.com/Balaganesh15M/demo-jenkin.git',
        credentialsId: 'your-git-credentials'  // Add your Git credentials ID here
      }
    }

    stage('Verify Setup') {
      steps {
        container('kaniko') {
          script {
            // Verify code was checked out
            sh 'ls -la /workspace'
            // Verify Docker config is mounted
            sh 'ls -la /kaniko/.docker/'
          }
        }
      }
    }

    stage('Build & Push Image') {
      steps {
        container('kaniko') {
          script {
            echo 'Building Docker image from Git repository...'
            // The actual build happens automatically via Kaniko's command
            // Verify Dockerfile exists
            sh 'test -f /workspace/Dockerfile && echo "Dockerfile found" || echo "Dockerfile missing"'
          }
        }
      }
    }
  }
}

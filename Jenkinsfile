pipeline {
  agent {
    kubernetes {
      defaultContainer 'jnlp'
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
    command: ["/kaniko/executor"]
    args: [
      "--dockerfile=/workspace/Dockerfile",
      "--context=/workspace",
      "--destination=docker.io/\${IMAGE}",
      "--verbosity=debug",
      "--skip-tls-verify"
    ]
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker/
    resources:
      limits:
        cpu: "1"
        memory: "1Gi"
      requests:
        cpu: "500m"
        memory: "512Mi"
  - name: jnlp
    image: jenkins/inbound-agent:latest
    resources:
      limits:
        cpu: "500m"
        memory: "512Mi"
      requests:
        cpu: "200m"
        memory: "256Mi"
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
        git branch: 'main',
            url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
      }
    }

    stage('Debug Workspace') {
      steps {
        container('kaniko') {
          sh 'echo "Verifying files in /workspace..."'
          sh 'ls -la /workspace'
          sh 'cat /workspace/Dockerfile || echo "Dockerfile not found!"'
        }
      }
    }

    stage('Build with Kaniko') {
      steps {
        container('kaniko') {
          echo "Kaniko build process should be running now..."
          // No additional commands needed; args do the build and push
        }
      }
    }
  }
}

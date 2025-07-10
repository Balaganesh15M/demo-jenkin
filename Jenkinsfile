pipeline {
  agent {
    kubernetes {
      defaultContainer 'jnlp'
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
    - --dockerfile=Dockerfile
    - --context=dir:///workspace
    - --destination=docker.io/bala1511/demo-jenkin:latest
    - --verbosity=debug
    - --skip-tls-verify
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker
    - name: workspace-volume
      mountPath: /workspace
    resources:
      limits:
        cpu: "200m"
        memory: "256Mi"
      requests:
        cpu: "100m"
        memory: "128Mi"
    securityContext:
      runAsUser: 0
  - name: jnlp
    image: jenkins/inbound-agent:latest
    volumeMounts:
    - name: workspace-volume
      mountPath: /workspace
    resources:
      limits:
        cpu: "250m"
        memory: "512Mi"
      requests:
        cpu: "100m"
        memory: "256Mi"
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
        credentialsId: 'github-creds'
      }
    }

    stage('Verify Setup') {
      steps {
        container('kaniko') {
          script {
            // Verify workspace content
            sh 'ls -la /workspace'
            // Verify Docker config
            sh 'ls -la /kaniko/.docker/'
            sh 'cat /kaniko/.docker/config.json || true'
            // Verify Dockerfile exists
            sh 'test -f /workspace/Dockerfile && echo "Dockerfile found" || echo "ERROR: Dockerfile missing"'
          }
        }
      }
    }

    stage('Build & Push Image') {
      steps {
        container('kaniko') {
          script {
            echo 'Starting Kaniko build process...'
            // The actual build happens automatically via the container command
          }
        }
      }
    }
  }
}

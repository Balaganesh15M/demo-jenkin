pipeline {
  agent {
    kubernetes {
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    jenkins-agent: kaniko
spec:
  containers:
    - name: kaniko
      image: gcr.io/kaniko-project/executor:v1.9.1
      command:
        - /kaniko/executor
      args:
        - --dockerfile=/workspace/source/Dockerfile
        - --context=dir:///workspace/source
        - --destination=docker.io/\${IMAGE}
        - --verbosity=debug
        - --skip-tls-verify
      volumeMounts:
        - name: docker-config
          mountPath: /kaniko/.docker/
        - name: workspace-volume
          mountPath: /workspace/source
      workingDir: /workspace/source
    - name: jnlp
      image: jenkins/inbound-agent:latest
      volumeMounts:
        - name: workspace-volume
          mountPath: /workspace/source
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

  environment {
    IMAGE = "bala1511/go-kaniko-demo:latest"
  }

  stages {
    stage('Checkout') {
      steps {
        container('jnlp') {
          dir('/workspace/source') {
            git branch: 'main',
                url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
          }
        }
      }
    }

    stage('Build & Push') {
      steps {
        container('kaniko') {
          echo "✅ Kaniko container will now build and push the image..."
          // No shell or sh step required
        }
      }
    }
  }
}

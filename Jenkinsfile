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
    command: ["/busybox/sh", "-c"]
    args: ["echo '== START Kaniko ==';
            ls -l /workspace/source;
            echo '== Docker Config ==';
            cat /kaniko/.docker/config.json || echo 'No config';
            echo '== Running Kaniko ==';
            /kaniko/executor --dockerfile=/workspace/source/Dockerfile --context=/workspace/source --destination=docker.io/\${IMAGE} --verbosity=debug --skip-tls-verify;
            echo '== Kaniko Finished =='"]
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker/
    - name: workspace-volume
      mountPath: /workspace/source
    workingDir: /workspace/source
    resources:
      limits:
        cpu: "1"
        memory: "1Gi"
      requests:
        cpu: "500m"
        memory: "512Mi"
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
    stage('Checkout Source') {
      steps {
        container('jnlp') {
          dir('/workspace/source') {
            git branch: 'main',
                url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
          }
        }
      }
    }

    stage('Debug Dockerfile Path') {
      steps {
        container('kaniko') {
          sh 'ls -l /workspace/source'
          sh 'cat /workspace/source/Dockerfile || echo "Dockerfile not found"'
        }
      }
    }

    stage('Build and Push Image') {
      steps {
        container('kaniko') {
          echo "Kaniko build and push starting..."
        }
      }
    }
  }
}

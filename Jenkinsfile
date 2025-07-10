pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    jenkins: agent
spec:
  containers:
  - name: golang
    image: golang:1.22
    command:
    - sleep
    args:
    - "9999"
    resources:
      requests:
        cpu: "100m"
        memory: "256Mi"
    volumeMounts:
    - mountPath: /home/jenkins/agent/workspace
      name: workspace-volume
    workingDir: /home/jenkins/agent/workspace

  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    command:
    - cat
    tty: true
    resources:
      requests:
        cpu: "100m"
        memory: "512Mi"
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker
    - name: workspace-volume
      mountPath: /home/jenkins/agent/workspace
    workingDir: /home/jenkins/agent/workspace

  - name: jnlp
    image: jenkins/inbound-agent:3309.v27b_9314fd1a_4-7
    resources:
      requests:
        cpu: "512m"
        memory: "512Mi"

  volumes:
  - name: docker-config
    secret:
      secretName: dockerhub-secret
      items:
      - key: .dockerconfigjson
        path: config.json
  - name: workspace-volume
    emptyDir: {}
"""
    }
  }

  environment {
    DOCKER_IMAGE = "docker.io/bala1511/userapi:latest"
    WORKDIR = "/home/jenkins/agent/workspace"
  }

  stages {

    stage('Checkout Code') {
      steps {
        git branch: 'main', url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
      }
    }

    stage('Verify Workspace') {
      steps {
        container('golang') {
          sh '''
            echo "=== Listing Workspace Contents ==="
            ls -l $WORKDIR
            echo "=== Dockerfile Content Preview ==="
            head -n 20 $WORKDIR/Dockerfile
          '''
        }
      }
    }

    stage('Build and Push Image with Kaniko') {
      steps {
        container('kaniko') {
          sh '''
            echo "=== Running Kaniko Build ==="
            /kaniko/executor \
              --context=dir://$WORKDIR \
              --dockerfile=$WORKDIR/Dockerfile \
              --destination=$DOCKER_IMAGE \
              --verbosity=debug
          '''
        }
      }
    }
  }

  post {
    failure {
      echo '❌ Build failed!'
    }
    success {
      echo '✅ Image built and pushed successfully.'
    }
  }
}

pipeline {
    agent {
        kubernetes {
            inheritFrom 'default'
            yaml """
kind: Pod
metadata:
  name: kaniko-build
spec:
  containers:
    - name: jnlp
      image: jenkins/inbound-agent:latest
      resources:
        requests:
          cpu: "100m"
          memory: "256Mi"

    - name: kaniko
      image: gcr.io/kaniko-project/executor:debug
      imagePullPolicy: Always
      command: ["cat"]
      tty: true
      volumeMounts:
        - mountPath: /kaniko/.docker
          name: docker-config
        - mountPath: /home/jenkins/agent/workspace
          name: workspace-volume
      workingDir: /home/jenkins/agent/workspace
      resources:
        requests:
          cpu: "100m"
          memory: "512Mi"

    - name: golang
      image: golang:1.22
      command: ["sleep", "9999"]
      tty: true
      volumeMounts:
        - mountPath: /home/jenkins/agent/workspace
          name: workspace-volume
      workingDir: /home/jenkins/agent/workspace
      resources:
        requests:
          cpu: "100m"
          memory: "256Mi"

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
        REPO_DIR = "/home/jenkins/agent/workspace"
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
            }
        }


      stage('Make Image') {
  steps {
    container('kaniko') {
      sh '''
        /kaniko/executor \
          --context=dir:///home/jenkins/agent/workspace/gggggg \
          --dockerfile=Dockerfile \
          --destination=docker.io/bala1511/userapi:latest \
          --verbosity=debug
      '''
    }
  }
}

    }
}

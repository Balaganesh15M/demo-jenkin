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
            environment {
                REGISTRY = 'index.docker.io'
                REPOSITORY = 'bala1511'
                IMAGE = 'jenkins-demo'
            }
            steps {
                container('kaniko') {
                    sh """
                    /kaniko/executor \
                      --dockerfile=${REPO_DIR}/Dockerfile.run \
                      --context=${REPO_DIR} \
                      --destination=${REGISTRY}/${REPOSITORY}/${IMAGE}:latest \
                      --cache=true \
                      --verbosity=debug
                    """
                }
            }
        }
    }
}

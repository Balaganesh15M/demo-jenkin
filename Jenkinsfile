pipeline {
    agent {
        kubernetes {
            inheritFrom 'default'  // Use your existing agent template
            yaml """
kind: Pod
metadata:
  name: kaniko
spec:
  containers:
  - name: golang
    image: golang:1.12
    command: ["cat"]
    tty: true
    volumeMounts:
      - name: workspace-volume
        mountPath: /workspace
    workingDir: /workspace
    resources:
      requests:
        cpu: "200m"
        memory: "256Mi"
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    imagePullPolicy: Always
    command: ["cat"]
    tty: true
    volumeMounts:
      - name: docker-config
        mountPath: /kaniko/.docker
      - name: workspace-volume
        mountPath: /workspace
    workingDir: /workspace
    resources:
      requests:
        cpu: "500m"
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
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                url: 'https://github.com/Balaganesh15M/demo-jenkin.git',
                credentialsId: 'github-creds'
            }
        }
        stage('Verify Setup') {
            steps {
                container('golang') {
                    sh 'ls -la /workspace'
                }
                container('kaniko') {
                    sh 'ls -la /kaniko/.docker'
                }
            }
        }
        stage('Build') {
            steps {
                container('golang') {
                    sh './build-go-bin.sh'
                }
            }
        }
        stage('Make Image') {
            environment {
                REGISTRY    = 'index.docker.io'
                REPOSITORY  = 'bala1511'
                IMAGE       = 'jenkins-demo'
            }
            steps {
                container('kaniko') {
                    script {
                        sh """
                        /kaniko/executor \
                        --dockerfile=Dockerfile.run \
                        --context=/workspace \
                        --destination=${REGISTRY}/${REPOSITORY}/${IMAGE}:latest \
                        --cache=true \
                        --verbosity=debug
                        """
                    }
                }
            }
        }
    }
}

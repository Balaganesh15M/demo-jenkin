pipeline {
    agent {
        kubernetes {
            label 'kaniko'
            yaml """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:latest
    command:
    - /busybox/sh
    args:
    - -c
    - sleep 3600
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker
    - name: workspace
      mountPath: /workspace
  volumes:
  - name: docker-config
    secret:
      secretName: dockerhub-secret
  - name: workspace
    emptyDir: {}
"""
        }
    }

    stages {
        stage('Checkout') {
            steps {
                container('kaniko') {
                    // Replace with your repository and branch details
                   git branch: 'main', 
                                      url: 'https://github.com/Balaganesh15M/demo-jenkin.git',
                }
            }
        }

        stage('Build and Push') {
            steps {
                container('kaniko') {
                    // Set the shell to busybox sh to avoid issues with other shells
                    withEnv(["PATH+MAVEN=/usr/local/maven/bin", "PATH=${WORKSPACE}/bin:${PATH}"]) {
                        sh """
                        /kaniko/executor \
                          --context=${WORKSPACE} \
                          --dockerfile=${WORKSPACE}/Dockerfile \
                          --destination=bala1511/jenkins-demo:latest
                        """
                    }
                }
            }
        }
    }
}

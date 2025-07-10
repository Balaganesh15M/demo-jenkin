pipeline {
    agent {
        kubernetes {
            inheritFrom 'default'
            yaml """
kind: Pod
metadata:
  name: kaniko
spec:
  containers:
  - name: golang
    image: golang:1.18
    command: ["cat"]
    tty: true
    volumeMounts:
    - mountPath: /home/jenkins/agent/workspace
      name: workspace-volume
    workingDir: /home/jenkins/agent/workspace
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
        
  - name: jnlp
    image: jenkins/inbound-agent:latest
    args: ['-url', 'http://jenkins.jenkins.svc.cluster.local:8080', 'SECRET', 'AGENT_NAME']
    volumeMounts:
    - mountPath: /home/jenkins/agent
      name: workspace-volume
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
        stage('Verify Setup') {
            steps {
                container('golang') {
                    script {
                        sh """
                        echo "Workspace contents:"
                        ls -la ${REPO_DIR}
                        echo "Build script exists?"
                        test -f ${REPO_DIR}/build-go-bin.sh && echo "Yes" || echo "No"
                        """
                    }
                }
            }
        }
        stage('Build') {
            steps {
                container('golang') {
                    script {
                        sh """
                        cd ${REPO_DIR}
                        chmod +x build-go-bin.sh
                        ./build-go-bin.sh
                        """
                    }
                }
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
                    script {
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
}

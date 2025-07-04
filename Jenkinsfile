pipeline {
    agent any

    environment {
        dockerimagename = "balaganesh15m/userapi:latest"
    }

    stages {
        stage('Checkout Source') {
            steps {
                git branch: 'main', url: 'https://github.com/Balaganesh15M/jenkins.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    dockerImage = docker.build(dockerimagename)
                }
            }
        }

        stage('Push Docker Image') {
            environment {
                registryCredentials = 'dockerhub'
            }
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', registryCredentials) {
                        dockerImage.push('latest')
                    }
                }
            }
        }

        stage('Deploy to Minikube') {
            steps {
                sh 'kubectl apply -f deployment.yaml'
            }
        }
    }
}

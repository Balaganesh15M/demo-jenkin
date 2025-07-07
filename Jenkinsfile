pipeline {
    agent any

    environment {
        dockerimagename = "bala1511/userapi:latest"     
    }

    stages {
        // stage('Checkout Source') {
        //     steps {
        //         git branch: 'main', url: 'https://github.com/Balaganesh15M/demo-jenkin.git'
        //     }
        // }

         stage('Checkout Source') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: 'main']],
                    userRemoteConfigs: [[
                        url: 'https://github.com/Balaganesh15M/demo-jenkin.git',
                        credentialsId: 'github-creds'
                    ]]
                ])
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
                sh 'kubectl rollout restart deployment/userapi'
            }
        }
    }
}

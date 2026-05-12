pipeline {
    agent any

    environment {
        DOCKER_USER = "sidra17"
        GIT_REPO_URL = "https://github.com/Shidqirasyad17/order-tariff-service.git" 
        IMAGE_NAME = "order-tariff-service"
    } 

    stages {
        stage('1. Checkout Code'){
            steps {
                git branch: 'main', url : "${GIT_REPO_URL}"
            }
        }

        stage('2. Unit Tests'){
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    sh 'go test ./internal/...'
                }
            }
        }

        stage('3. Lint/Vet'){
            steps {
                sh 'go vet ./...'
            }
        }

        stage('4. Build Docker Image'){
            steps {

                sh "docker build -t ${IMAGE_NAME}:latest ."
            }
        }

        stage('5. Functional Test'){
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    sh 'docker-compose up -d postgres redis'
                    sh 'sleep 10'
                    sh 'go test ./test/...'
                    sh 'docker-compose down'
                }
            }
        }

        stage('6. Push image'){
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-login', passwordVariable :'PASS', usernameVariable: 'USER')]){
                    sh "echo ${PASS} | docker login -u ${USER} --password-stdin"
                    sh "docker tag ${IMAGE_NAME}:latest ${USER}/${IMAGE_NAME}:latest"
                    sh "docker push ${USER}/${IMAGE_NAME}:latest"
                }
            }
        }

        stage('7. Deploy to Kubernetes') {
            steps {
                echo 'Deploying image ke cluster Kubernetes lokal (Docker Desktop)...'
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
    
                    sh "kubectl apply -f deployment.yaml"
                }
            }
        }

        stage('8. Verify') {
            steps {
                echo 'Verifikasi status deployment...'
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    sh "kubectl get pods"
                    sh "kubectl rollout status deployment/order-tariff-deployment"
                }
            }
        }
    }
}
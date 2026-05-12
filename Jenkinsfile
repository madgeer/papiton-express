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
                    bat 'go test ./internal/...'
                }
            }
        }

        stage('3. Lint/Vet'){
            steps {
                bat 'go vet ./...'
            }
        }

        stage('4. Build Docker Image'){
            steps {
               
                bat "docker build -t %IMAGE_NAME%:latest ."
            }
        }

        stage('5. Functional Test'){
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    bat 'docker-compose up -d postgres redis'
        
                    bat 'timeout /t 10 /nobreak'
                    bat 'go test ./test/...'
                    bat 'docker-compose down'
                }
            }
        }

        stage('6. Push image'){
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-login', passwordVariable :'PASS', usernameVariable: 'USER')]){
                    
                    bat "echo %PASS% | docker login -u %USER% --password-stdin"
                    bat "docker tag %IMAGE_NAME%:latest %USER%/%IMAGE_NAME%:latest"
                    bat "docker push %USER%/%IMAGE_NAME%:latest"
                }
            }
        }

        stage('7. Deploy to Kubernetes') {
            steps {
                echo 'Deploying image ke cluster Kubernetes lokal (Docker Desktop)...'
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    bat "kubectl apply -f deployment.yaml"
                }
            }
        }

        stage('8. Verify') {
            steps {
                echo 'Verifikasi status deployment...'
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    bat "kubectl get pods"
                    bat "kubectl rollout status deployment/order-tariff-deployment"
                }
            }
        }
    }
}
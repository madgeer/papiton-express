pipeline {
    agent any

    environment {
        DOCKER_USER = "madgeer"
        GIT_REPO_URL = "https://github.com/Shidqirasyad17/order-tariff-service.git" // Default git URL dari project asli
    } 

    stages {
        // =====================================================================
        // STAGE 1: CHECKOUT REPO & DETECT CHANGES
        // =====================================================================
        stage('1. Checkout Repo & Detect Changes') {
            steps {
                script {
                    // Deteksi berkas yang berubah pada commit terakhir
                    // Menggunakan bat karena agent Jenkins berjalan di OS Windows
                    def changedFiles = bat(
                        script: "@echo off\ngit diff --name-only HEAD~1 HEAD",
                        returnStdout: true
                    ).trim().split("\r?\n")
                    
                    env.CHANGE_WAREHOUSE = "false"
                    env.CHANGE_ORDER = "false"
                    env.CHANGE_SHIPPING = "false"

                    for (file in changedFiles) {
                        if (file.startsWith("warehouse-and-inventory-service/")) {
                            env.CHANGE_WAREHOUSE = "true"
                        }
                        if (file.startsWith("order-tariff-service/")) {
                            env.CHANGE_ORDER = "true"
                        }
                        if (file.startsWith("shipping-service/")) {
                            env.CHANGE_SHIPPING = "true"
                        }
                    }
                    
                    echo "--- DETEKSI PERUBAHAN SERVICE ---"
                    echo "Warehouse Service: ${env.CHANGE_WAREHOUSE}"
                    echo "Order-Tariff Service: ${env.CHANGE_ORDER}"
                    echo "Shipping Service: ${env.CHANGE_SHIPPING}"
                    echo "---------------------------------"
                }
            }
        }

        // =====================================================================
        // STAGE 2: UNIT TESTS (go test)
        // =====================================================================
        stage('2. Unit Tests') {
            steps {
                script {
                    if (env.CHANGE_WAREHOUSE == "true") {
                        dir('warehouse-and-inventory-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat 'go test ./internal/...'
                            }
                        }
                    }
                    if (env.CHANGE_ORDER == "true") {
                        dir('order-tariff-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat 'go test ./internal/...'
                            }
                        }
                    }
                }
            }
        }

        // =====================================================================
        // STAGE 3: LINT/VET (go vet)
        // =====================================================================
        stage('3. Lint/Vet') {
            steps {
                script {
                    if (env.CHANGE_WAREHOUSE == "true") {
                        dir('warehouse-and-inventory-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat 'go vet ./...'
                            }
                        }
                    }
                    if (env.CHANGE_ORDER == "true") {
                        dir('order-tariff-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat 'go vet ./...'
                            }
                        }
                    }
                }
            }
        }

        // =====================================================================
        // STAGE 4: BUILD IMAGE (lokal)
        // =====================================================================
        stage('4. Build Image') {
            steps {
                script {
                    if (env.CHANGE_WAREHOUSE == "true") {
                        bat "docker build -t %DOCKER_USER%/papiton-warehouse-service:latest ./warehouse-and-inventory-service"
                    }
                    if (env.CHANGE_ORDER == "true") {
                        // order-tariff menggunakan lowercase 'dockerfile'
                        bat "docker build -t %DOCKER_USER%/order-tariff-service:latest ./order-tariff-service -f ./order-tariff-service/dockerfile"
                    }
                    if (env.CHANGE_SHIPPING == "true") {
                        bat "docker build -t %DOCKER_USER%/shipping-service:latest ./shipping-service"
                    }
                }
            }
        }

        // =====================================================================
        // STAGE 5: FUNCTIONAL TESTS (test app di lokal atau staging)
        // =====================================================================
        stage('5. Functional Tests') {
            steps {
                script {
                    // Functional Test untuk Warehouse (menggunakan Testcontainers)
                    if (env.CHANGE_WAREHOUSE == "true") {
                        dir('warehouse-and-inventory-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat 'go test ./test/functional/...'
                            }
                        }
                    }
                    
                    // Functional Test untuk Order-Tariff (menggunakan docker-compose dari root)
                    if (env.CHANGE_ORDER == "true") {
                        catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                            bat 'docker-compose -f docker-compose.yml up -d order-db order-redis'
                            bat 'ping 127.0.0.1 -n 11 > nul' // Menunggu db siap
                            
                            withEnv(["DB_PORT=5434", "DB_HOST=localhost", "DB_USER=postgres", "DB_PASSWORD=admin123", "DB_NAME=papiton_order_tariff_service_db"]) {
                                dir('order-tariff-service') {
                                    bat 'go test ./tests/...'
                                }
                            }
                            
                            bat 'docker-compose -f docker-compose.yml down'
                        }
                    }
                }
            }
        }

        // =====================================================================
        // STAGE 6: PUSH IMAGE
        // =====================================================================
        stage('6. Push image') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'dockerhub-login', passwordVariable :'PASS', usernameVariable: 'USER')]){
                        catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                            bat "docker login -u %USER% -p %PASS%"
                            
                            if (env.CHANGE_WAREHOUSE == "true") {
                                bat "docker tag %DOCKER_USER%/papiton-warehouse-service:latest %USER%/papiton-warehouse-service:latest"
                                bat "docker push %USER%/papiton-warehouse-service:latest"
                            }
                            if (env.CHANGE_ORDER == "true") {
                                bat "docker tag %DOCKER_USER%/order-tariff-service:latest %USER%/order-tariff-service:latest"
                                bat "docker push %USER%/order-tariff-service:latest"
                            }
                            if (env.CHANGE_SHIPPING == "true") {
                                bat "docker tag %DOCKER_USER%/shipping-service:latest %USER%/shipping-service:latest"
                                bat "docker push %USER%/shipping-service:latest"
                            }
                        }
                    }
                }
            }
        }

        // =====================================================================
        // STAGE 7: DEPLOY DI KUBERNETES
        // =====================================================================
        stage('7. Deploy di kubernetes') {
            steps {
                script {
                    if (env.CHANGE_WAREHOUSE == "true") {
                        dir('warehouse-and-inventory-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat "kubectl apply -f deployment.yaml"
                            }
                        }
                    }
                    if (env.CHANGE_ORDER == "true") {
                        dir('order-tariff-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat "kubectl apply -f deployment.yaml"
                            }
                        }
                    }
                    if (env.CHANGE_SHIPPING == "true") {
                        dir('shipping-service') {
                            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                                bat "kubectl apply -f k8s/deployment.yaml"
                            }
                        }
                    }
                }
            }
        }

        // =====================================================================
        // STAGE 8: VERIFY
        // =====================================================================
        stage('8. Verify') {
            steps {
                script {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        bat "kubectl get pods"
                        
                        if (env.CHANGE_WAREHOUSE == "true") {
                            bat "kubectl rollout status deployment/warehouse-deployment"
                        }
                        if (env.CHANGE_ORDER == "true") {
                            bat "kubectl rollout status deployment/order-tariff-deployment"
                        }
                        if (env.CHANGE_SHIPPING == "true") {
                            bat "kubectl rollout status deployment/shipping-deployment"
                        }
                    }
                }
            }
        }
    }
}

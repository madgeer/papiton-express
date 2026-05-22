def getServices() {
    return [
        [
            id: 'order-tariff',
            path: 'order-tariff-service',
            dockerfile: 'dockerfile',
            deployment: 'order-tariff-deployment',
            container: 'order-tariff-container',
            imageName: 'order-tariff-service',
            functionalTestPath: './tests/...'
        ],
        [
            id: 'shipping',
            path: 'shipping-service',
            dockerfile: 'Dockerfile',
            deployment: 'shipping-service-deployment',
            container: 'shipping-service',
            imageName: 'shipping-service',
            functionalTestPath: './internal/repository/... ./internal/handler/...'
        ],
        [
            id: 'warehouse',
            path: 'warehouse-and-inventory-service',
            dockerfile: 'Dockerfile',
            deployment: 'warehouse-deployment',
            container: 'warehouse-container',
            imageName: 'warehouse-service',
            functionalTestPath: './test/functional/...'
        ],
        [
            id: 'tracking',
            path: 'tracking-service-and-logevent-service',
            dockerfile: 'Dockerfile',
            deployment: 'tracking-service',
            container: 'tracking-service',
            imageName: 'tracking-service',
            functionalTestPath: './internal/service/...'
        ],
        [
            id: 'notification', 
            path: 'notification-and-messaging-service',
            dockerfile: 'Dockerfile',
            deployment: 'notification-service',
            container: 'notification-service',
            imageName: 'notification-service',
            functionalTestPath: './tests/functional/...'
        ]
    ]
}

node {
    def registry = "madgeer" // Ganti dengan URL registry Docker Anda
    def gitBranch = 'main' // Ganti dengan branch target Anda jika berbeda
    def gitUrl = 'https://github.com/madgeer/papiton-express.git'

    // Menyimpan status kegagalan untuk setiap tahapan verifikasi
    def stageFailures = [:]

    try {
        stage('Checkout') {
            echo "Checking out repository from branch ${gitBranch}..."
            git url: gitUrl, branch: gitBranch
        }

        // --- STAGE: UNIT TESTS ---
        try {
            stage('Unit Tests') {
                docker.image('golang:1.26-alpine').inside {
                    def services = getServices()
                    def failedServices = []
                    services.each { service ->
                        echo "Running unit tests for ${service.id}..."
                        try {
                            dir(service.path) {
                                sh 'go test -short ./...'
                            }
                        } catch (Exception e) {
                            echo "Unit tests failed for ${service.id}!"
                            failedServices.add(service.id)
                        }
                    }
                    if (failedServices.size() > 0) {
                        error "Unit tests failed for: ${failedServices}"
                    }
                }
            }
        } catch (Exception e) {
            stageFailures['Unit Tests'] = e.message
            currentBuild.result = 'UNSTABLE'
        }

        // --- STAGE: LINT/VET ---
        try {
            stage('Lint/Vet') {
                docker.image('golang:1.26-alpine').inside {
                    def services = getServices()
                    def failedServices = []
                    services.each { service ->
                        echo "Running go vet for ${service.id}..."
                        try {
                            dir(service.path) {
                                sh 'go vet ./...'
                            }
                        } catch (Exception e) {
                            echo "Go vet failed for ${service.id}!"
                            failedServices.add(service.id)
                        }
                    }
                    if (failedServices.size() > 0) {
                        error "Go vet failed for: ${failedServices}"
                    }
                }
            }
        } catch (Exception e) {
            stageFailures['Lint/Vet'] = e.message
            currentBuild.result = 'UNSTABLE'
        }

        // --- STAGE: BUILD IMAGE ---
        try {
            stage('Build Image') {
                def services = getServices()
                def failedServices = []
                services.each { service ->
                    echo "Building Docker image for ${service.id}..."
                    try {
                        dir(service.path) {
                            sh "docker build -f ${service.dockerfile} -t ${service.imageName}:${BUILD_NUMBER} ."
                        }
                    } catch (Exception e) {
                        echo "Docker build failed for ${service.id}!"
                        failedServices.add(service.id)
                    }
                }
                if (failedServices.size() > 0) {
                    error "Docker build failed for: ${failedServices}"
                }
            }
        } catch (Exception e) {
            stageFailures['Build Image'] = e.message
            currentBuild.result = 'UNSTABLE'
        }

        // --- STAGE: START DATABASES (Untuk Functional Tests) ---
        try {
            stage('Start Databases') {
                echo 'Starting database containers for functional tests...'
                sh 'docker rm -f shipping-db shipping-mongo order-db order-redis || true'
                sh 'docker run -d --name shipping-db -p 5433:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=shipping_test_db postgres:15-alpine'
                sh 'docker run -d --name shipping-mongo -p 27017:27017 mongo:6-jammy'
                sh 'docker run -d --name order-db -p 5434:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin123 -e POSTGRES_DB=papiton_order_tariff_service_db postgres:15-alpine'
                sh 'docker run -d --name order-redis -p 6379:6379 redis:alpine'
                echo 'Waiting for databases to be ready...'
                sh 'sleep 10'
            }
        } catch (Exception e) {
            stageFailures['Start Databases'] = e.message
            currentBuild.result = 'UNSTABLE'
        }

        // --- STAGE: FUNCTIONAL TESTS ---
        try {
            stage('Functional Tests') {
                docker.image('golang:1.26-alpine').inside('--network host -v /var/run/docker.sock:/var/run/docker.sock') {
                    def services = getServices()
                    def failedServices = []
                    services.each { service ->
                        echo "Running functional tests for ${service.id}..."
                        try {
                            dir(service.path) {
                                sh "go test -tags functional -v ${service.functionalTestPath}"
                            }
                        } catch (Exception e) {
                            echo "Functional tests failed for ${service.id}!"
                            failedServices.add(service.id)
                        }
                    }
                    if (failedServices.size() > 0) {
                        error "Functional tests failed for: ${failedServices}"
                    }
                }
            }
        } catch (Exception e) {
            stageFailures['Functional Tests'] = e.message
            currentBuild.result = 'UNSTABLE'
        }

        // --- CD STAGES: HANYA BERJALAN JIKA SELURUH VERIFIKASI BERHASIL ---
        if (stageFailures.size() == 0) {
            stage('Push Image') {
                def services = getServices()
                services.each { service ->
                    echo "Pushing image for ${service.id} to registry..."
                    sh "docker tag ${service.imageName}:${BUILD_NUMBER} ${registry}/${service.imageName}:${BUILD_NUMBER}"
                    sh "docker push ${registry}/${service.imageName}:${BUILD_NUMBER}"
                }
            }

            stage('Deploy to Kubernetes') {
                def services = getServices()
                services.each { service ->
                    echo "Deploying ${service.id} to Kubernetes..."
                    sh "kubectl set image deployment/${service.deployment} ${service.container}=${registry}/${service.imageName}:${BUILD_NUMBER} -n production"
                }
            }

            stage('Verify') {
                def services = getServices()
                services.each { service ->
                    echo "Verifying deployment for ${service.id}..."
                    sh "kubectl rollout status deployment/${service.deployment} -n production"
                }
            }

            currentBuild.result = 'SUCCESS'
            echo 'Pipeline completed successfully!'
        } else {
            // Tandai build sebagai gagal di akhir jika ada tahapan verifikasi yang gagal
            currentBuild.result = 'FAILURE'
            echo "Pipeline completed with failures in: ${stageFailures.keySet()}"
            error "Pipeline failed due to stage failures: ${stageFailures}"
        }

    } catch (err) {
        currentBuild.result = 'FAILURE'
        echo "Pipeline aborted due to critical error: ${err.message}"
        throw err
    } finally {
        stage('Cleanup') {
            echo 'Cleaning up database containers...'
            sh 'docker rm -f shipping-db shipping-mongo order-db order-redis || true'
            echo 'Pipeline execution completed.'
        }
    }
}

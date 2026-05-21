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
        ]
    ]
}

pipeline {
    agent any

    environment {
        REGISTRY = "<registry>" // Ganti dengan URL registry Docker Anda
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Checking out repository...'
                checkout scm
            }
        }

        stage('Unit Tests') {
            steps {
                script {
                    def services = getServices()
                    services.each { service ->
                        echo "Running unit tests for ${service.id}..."
                        dir(service.path) {
                            sh 'go test -short ./...'
                        }
                    }
                }
            }
        }

        stage('Lint/Vet') {
            steps {
                script {
                    def services = getServices()
                    services.each { service ->
                        echo "Running go vet for ${service.id}..."
                        dir(service.path) {
                            sh 'go vet ./...'
                        }
                    }
                }
            }
        }

        stage('Build Image') {
            steps {
                script {
                    def services = getServices()
                    services.each { service ->
                        echo "Building Docker image for ${service.id}..."
                        dir(service.path) {
                            sh "docker build -f ${service.dockerfile} -t ${service.imageName}:${BUILD_NUMBER} ."
                        }
                    }
                }
            }
        }

        stage('Start Databases') {
            steps {
                script {
                    echo 'Starting database containers for functional tests...'
                    sh 'docker compose up -d order-db order-redis shipping-db shipping-mongo warehouse-db'
                    echo 'Waiting for databases to be ready...'
                    sh 'sleep 10'
                }
            }
        }

        stage('Functional Tests') {
            steps {
                script {
                    def services = getServices()
                    services.each { service ->
                        echo "Running functional tests for ${service.id}..."
                        dir(service.path) {
                            sh "go test -tags functional -v ${service.functionalTestPath}"
                        }
                    }
                }
            }
        }

        stage('Push Image') {
            steps {
                script {
                    def services = getServices()
                    services.each { service ->
                        echo "Pushing image for ${service.id} to registry..."
                        sh "docker tag ${service.imageName}:${BUILD_NUMBER} ${REGISTRY}/${service.imageName}:${BUILD_NUMBER}"
                        sh "docker push ${REGISTRY}/${service.imageName}:${BUILD_NUMBER}"
                    }
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                script {
                    def services = getServices()
                    services.each { service ->
                        echo "Deploying ${service.id} to Kubernetes..."
                        sh "kubectl set image deployment/${service.deployment} ${service.container}=${REGISTRY}/${service.imageName}:${BUILD_NUMBER} -n production"
                    }
                }
            }
        }

        stage('Verify') {
            steps {
                script {
                    def services = getServices()
                    services.each { service ->
                        echo "Verifying deployment for ${service.id}..."
                        sh "kubectl rollout status deployment/${service.deployment} -n production"
                    }
                }
            }
        }
    }

    post {
        always {
            echo 'Cleaning up database containers...'
            sh 'docker compose down || true'
            echo 'Pipeline execution completed.'
        }
        failure {
            echo 'Pipeline failed. Please check logs.'
        }
        success {
            echo 'Pipeline succeeded.'
        }
    }
}

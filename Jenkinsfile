pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {

                git url: 'https://github.com/c0ldw4ter/jenkins-docker.git', branch: 'main'
            }
        }

        stage('Build Image') {
            steps {

                sh '''
                  echo "Stopping old container if exists..."
                  docker stop test_jenkins_nginx || true
                  sleep 5

                  echo "BUILDING DOCKER"
                  docker build -t trainee-jenkins .
                '''
            }
        }

        stage('Test Container') {
            steps {
                sh '''
                  echo "RUNNING DOCKER"
                  docker run --rm -d --name=test_jenkins_nginx -p 8888:80 trainee-jenkins

                  sleep 5
                  echo "START TEST"
                  curl http://158.160.75.45:8888 || (echo "curl failed!" && exit 1)

   
                  docker stop test_jenkins_nginx || true
                '''
            }
        }

        stage('Push to Docker Hub') {
            steps {
            
                    sh '''
                      echo "TO REGISTRY"

                      docker login --username "$DOCKER_USER" --password "$DOCKER_PASS"

                      docker tag trainee-jenkins:latest c0ldw4ter/jenkis-docker:${BUILD_ID}
                      docker push c0ldw4ter/jenkis-docker:${BUILD_ID}
                    '''
                
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                  echo "DEPLOY"
                  ansible-playbook /etc/ansible/pipe-dep.yaml --extra-vars TAG=${BUILD_ID}
                '''
            }
        }
    }
}

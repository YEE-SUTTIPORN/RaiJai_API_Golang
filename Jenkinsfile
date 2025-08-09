pipeline {
  agent any

  environment {
    IMAGE_NAME = 'ayeesuttiporn/raijai_api_repo'
    IMAGE_TAG = "${env.BUILD_NUMBER}"
    DEPLOY_REPO = 'git@github.com:your-org/k8s-manifests.git'
  }

  stages {
    stage('Checkout') {
      steps { git 'https://github.com/YOUR_NAME/raijai_project.git' }
    }

    stage('Build Docker') {
      steps { sh "docker build -t $IMAGE_NAME:$IMAGE_TAG ." }
    }

    stage('Push Docker') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
          sh "echo $PASS | docker login -u $USER --password-stdin"
          sh "docker push $IMAGE_NAME:$IMAGE_TAG"
        }
      }
    }

    stage('Update Deployment YAML') {
      steps {
        dir('k8s') {
          git credentialsId: 'your-deploy-key', url: "$DEPLOY_REPO"
          sh "sed -i 's|image: .*|image: $IMAGE_NAME:$IMAGE_TAG|' raijai-api/deployment.yaml"
          sh "git config user.email 'ci@your.com'"
          sh "git config user.name 'CI Bot'"
          sh "git commit -am 'update image to $IMAGE_TAG'"
          sh "git push origin main"
        }
      }
    }
  }
}

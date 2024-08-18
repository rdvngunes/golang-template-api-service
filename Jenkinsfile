pipeline {
  environment {
    registry = "golangtemplates/golang-template-api-service"
    registryCredential = 'golangtemplates'
    dockerImage = ''
    buildVersion = "$BRANCH_NAME-V$BUILD_NUMBER".replaceAll('/','-')
    latestVersion = "$BRANCH_NAME-latest".replaceAll('/','-')
  }
  parameters {
    //Jenkins jobs execution parameters
    string(name: 'sshServer', defaultValue: '[[user@api]]', description: 'Deployment Server (ubuntu@xx.xx.xx.xx)')
	}
  agent any
  stages {
    stage('Building Image')
    {
    steps {
        sh 'printenv'
           script {
                dockerImage = docker.build registry + ":$buildVersion"
            }
        }
    }
    stage('Push Image to DockHub') {
      steps{
        script {
          docker.withRegistry( '', registryCredential ) {
            dockerImage.push("$buildVersion")
            dockerImage.push("$latestVersion")
          }
        }
      }
    }
    stage('Remove Unused docker image') {
      steps{
        sh "docker rmi $registry:$buildVersion"
        sh "docker rmi $registry:$latestVersion"
      }
    }
	stage('Run docker container') {
    
      steps{
        sshagent (credentials: ['soramate']){
          sh "ssh -o StrictHostKeyChecking=no -T ${params.sshServer} sudo docker kill \"golang-template-api-service\""
          sh "ssh -o StrictHostKeyChecking=no -T ${params.sshServer} sudo docker rm \"golang-template-api-service\""
          sh "ssh -o StrictHostKeyChecking=no -T ${params.sshServer} sudo docker run -d --name \"golang-template-api-service\" -p 3000:3000 --restart unless-stopped $registry:$buildVersion"
        }
      }
    }
  }
}
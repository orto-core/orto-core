pipeline {

  agent any
  stages{
    stage("Test"){
      steps {
        echo "Hello World!"
      }
    }  
  }
  post {
    always {
      echo 'One way or another, I have finished'
      cleanWs()
    }

    success {
      script {
        echo "Everything was successful"
      }
    }

    unstable {
      echo 'I am unstable :/'
    }

    failure {
      script {
        echo "Failure everywhere"
      }
    }

    changed {
      echo 'Things were different before...'
    	}
  }
}
 

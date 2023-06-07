#!groovy
println('---------------------------------------------------------------------------Import Job CI/Dviross')
// define first pipeline 
def pipelineScript = new File('/var/jenkins_config/jobs/Dviross_wrapper.groovy').getText("UTF-8")
//Adding pipeline to the right path 
pipelineJob('CI/Dviross_wrapper'){
    description("Build .jar file from Dviross repository , from master branch")
//Importing parameter 
    parameters {
        stringParam {
            name('BRANCH')
            defaultValue('master')
            description("branch to pull")
            trim(false)
        }
        booleanParam {
            name('SKIP_TEST')
            defaultValue(true)
            description("Skip test")
        }
        choice {
            name('VERSION_TYPE')
            choices(['SNAPSHOT', 'RELEASE'])
            description('Version type between snapshot and release')
        }
        stringParam {
            name('VERSION')
            defaultValue('1.0')
            description("version of the project")
            trim(false)
        }
    }
//Importing the script in sandbox 
    definition{
        cps {
            script(pipelineScript)
            sandbox()
        }
    }
}
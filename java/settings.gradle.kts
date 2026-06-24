pluginManagement {
  repositories {
    gradlePluginPortal()
    mavenCentral()
  }
}

dependencyResolutionManagement {
  repositoriesMode.set(RepositoriesMode.FAIL_ON_PROJECT_REPOS)
  repositories {
    mavenCentral()
    mavenLocal()
  }
}

rootProject.name = "runapi-gpt-4o-image-java"

include("runapi-core")
include("runapi-gpt-4o-image")

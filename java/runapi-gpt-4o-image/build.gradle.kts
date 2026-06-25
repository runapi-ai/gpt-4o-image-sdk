plugins {
  `java-library`
  `maven-publish`
}

description = "RunAPI GPT-4o Image Java SDK for GPT-4o Image workflows."

java {
  withSourcesJar()
  withJavadocJar()
}

dependencies {
  api("ai.runapi:runapi-core:0.1.1")

  testImplementation(platform("org.junit:junit-bom:5.10.3"))
  testImplementation("org.junit.jupiter:junit-jupiter")
}

publishing {
  publications {
    create<MavenPublication>("mavenJava") {
      from(components["java"])
      artifactId = "runapi-gpt-4o-image"
      pom {
        name = "RunAPI GPT-4o Image Java SDK"
        description = "RunAPI GPT-4o Image Java SDK for GPT-4o Image workflows."
        url = "https://runapi.ai/models/gpt-4o-image"
        licenses {
          license {
            name = "Apache License, Version 2.0"
            url = "https://www.apache.org/licenses/LICENSE-2.0"
          }
        }
        developers {
          developer {
            id = "runapi"
            name = "RunAPI"
            email = "contact@runapi.ai"
          }
        }
        scm {
          url = "https://github.com/runapi-ai/gpt-4o-image-sdk"
          connection = "scm:git:https://github.com/runapi-ai/gpt-4o-image-sdk.git"
          developerConnection = "scm:git:ssh://git@github.com/runapi-ai/gpt-4o-image-sdk.git"
        }
      }
    }
  }
}

variable "GO_RACE_DETECTOR" {
}

variable "TAG" {
  default = "latest"
}

function "getTags" {
  params = [imageName]
  result = [
    "${imageName}:${TAG}",
  ]
}

group "default" {
  targets = [
    "task-management-api",
    "task-management-migration-db",
    "task-management-api-docs"
  ]
}

target "builder" {
  dockerfile = "docker/builder/Dockerfile"
  target = "builder"
  args = {
    GO_RACE_DETECTOR = notequal("",GO_RACE_DETECTOR) ? "true" : "false"
  }
}

target "runner" {
  dockerfile = "docker/builder/Dockerfile"
  target = "runner"
}

target "service-base" {
  contexts = {
    builder = "target:builder"
    runner = "target:runner"
  }
}

target "task-management-api" {
  inherits = ["service-base"]
  dockerfile = "docker/api/Dockerfile"
  tags = getTags("task-management-api")
}

target "task-management-migration-db" {
  dockerfile = "docker/migration-db/Dockerfile"
  tags = getTags("task-management-migration-db")
}

target "task-management-api-docs" {
  dockerfile = "docker/api-docs/Dockerfile"
  tags = getTags("task-management-api-docs")
}

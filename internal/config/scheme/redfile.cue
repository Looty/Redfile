#Version: {
	version: number | *1.0
}

#Env: {
  [string]: string
}

#Task: {
  [string]: {
    cmds: [
      ...string
    ]
    env?: {
      #Env
    }
    description?: {
      string
    }
  }
}

#Version
env: {
  {...,#Env}
}
tasks: {
  {...,#Task}
}
import { TodoSchema } from "./types";


export const getTodos = async () => {
  try {
    const resp = await fetch(import.meta.env.VITE_TODO_API, {
      method: "GET",
    })
    const data = await resp.json()
    if (!resp.ok) {
      throw Error(data.msg)
    }
    return data
  } catch (error) {
    throw Error(`${error}`)
  }
}

export const getTodo = async (id: string) => {
  try {
    const resp = await fetch(import.meta.env.VITE_TODO_API + `/${encodeURI(id)}`, {
      method: "GET",
    })
    const data = await resp.json()
    if (!resp.ok || data===null) {
      throw Error(data.msg)
    }
    return data
  } catch (error) {
    throw Error(`${error}`)
  }
}

export const createTodo = async (todo: TodoSchema) => {
  try {
    const resp = await fetch(import.meta.env.VITE_TODO_API, {
      method: "POST",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(todo)
    })
    const data = await resp.json()
    if (!resp.ok) {
      throw Error(data.msg)
    }
    return data
  } catch (error) {
    throw Error(`${error}`)
  }
}

export const toggleTodo = async (todo: TodoSchema) => {
  try {
    const resp = await fetch(import.meta.env.VITE_TODO_API, {
      method: "PUT",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(todo)
    })
    const data = await resp.json()
    if (!resp.ok) {
      throw Error(data.msg)
    }
    return data
  } catch (error) {
    throw Error(`${error}`)
  }
}


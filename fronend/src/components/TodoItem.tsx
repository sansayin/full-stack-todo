import { useQueryClient, useMutation } from '@tanstack/react-query';
import { toggleTodo } from '../lib/todo-api';
import { TodoSchema } from '../lib/types';
import _ from 'lodash';
import { toast } from 'react-toastify';

type TodoItemProps = {
  todo: TodoSchema;
  id: number
};

export const TodoItem = (props: TodoItemProps) => {

  const queryClient = useQueryClient();

  const updateCache = (newTodo: TodoSchema) => {
    queryClient.setQueryData(['todos'], (old: TodoSchema[]) => {
      const newTodos = old.map(todo => 
        todo= todo.id == newTodo.id? newTodo:todo
      )
      console.log(JSON.stringify(newTodos))
      return [...newTodos]
    })
  }

  const todoMutation = useMutation({
    mutationFn: toggleTodo,
    onMutate: async (newTodo: TodoSchema) => {
      await queryClient.cancelQueries(['todos'])
      const prevTodos = queryClient.getQueryData<TodoSchema[]>(['todos'])
      updateCache(newTodo)
      return { prevTodos } //context on onError
    },
    onSuccess: ({ data }) => {
      toast(`${data.description} is ${data.completed?"completed":"in todo list"}`);
    },
    onError: (err, _, context) => {
      toast(err.message);
      queryClient.setQueryData(["todos"], context?.prevTodos)
    },
  });

  const toggleTodoCompletion = async () => {
    await todoMutation.mutateAsync({ ...props.todo, completed: !props.todo.completed });
  };
  return (
    <li>
      <input
        type="checkbox"
        checked={props.todo.completed}
        onChange={toggleTodoCompletion}
      />
      <span>{props.todo.description}</span>
    </li>
  );
};


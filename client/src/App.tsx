import { Button, Container, Stack } from '@chakra-ui/react'
import './App.css'
import NavBar from './components/NavBar';
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'
import React from 'react';

export const BASE_URL = "/api";

function App() {
  return (
    <Stack h="100vh">
      <NavBar />

      <Container>
        <TodoForm />
        <TodoList />
      </Container>
    </Stack>
  );
}

export default App;

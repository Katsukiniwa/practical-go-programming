CREATE USER 'user-katsukiniwa' @'%' IDENTIFIED BY 'password';
GRANT SELECT,
  INSERT,
  UPDATE,
  DELETE ON todo.* TO 'user-katsukiniwa' @'%';

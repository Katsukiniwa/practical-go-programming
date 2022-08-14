CREATE USER 'user-katsukiniwa' @'%' IDENTIFIED BY 'katsukiniwa-password';
GRANT SELECT,
  INSERT,
  UPDATE,
  DELETE ON todo.* TO 'user-katsukiniwa' @'%';

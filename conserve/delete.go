package conserve

func Delete(executable, dir, name string) error {
	_, err := executeDelete(executable, dir, name)
	return err
}

// Delete executes the conserve delete command.
// It returns the output of the command can be used for logging and an error if any.
func executeDelete(executable, dir, name string) (string, error) {
	return execute(executable, "delete", "--backup", name, dir)
}

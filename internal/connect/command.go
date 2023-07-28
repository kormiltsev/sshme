package connect

// SetCommand just put command in Job structure.
func (j *Job) SetCommand(s string) error {
	j.Command = s
	return nil
}

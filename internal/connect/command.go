package connect

func (j *Job) SetCommand(s string) error {
	j.Command = s
	return nil
}

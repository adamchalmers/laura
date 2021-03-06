/*
 * An implementation of the laura/filesys.FileSys interface.
 * Used for testing app logic against a fake instead of a real filesystem.
 */

package filesys

type fakeFS struct {
	names []string
	diary map[string]string
}

func NewFakeFS() *fakeFS {
	fs := new(fakeFS)
	fs.diary = make(map[string]string, 0)
	return fs
}

func (fs *fakeFS) Names() []string {
	return fs.names
}
func (fs *fakeFS) ReadDiary(name string) (string, error) {
	return fs.diary[name], nil
}
func (fs *fakeFS) MakeDiary(name string) error {
	fs.names = append(fs.names, name)
	fs.diary[name] = ""
	return nil
}
func (fs *fakeFS) DeleteDiary(target string) error {
	newNames := make([]string, 0)
	for _, name := range fs.names {
		if name != target {
			newNames = append(newNames, target)
		}
	}
	fs.names = newNames
	return nil
}
func (fs *fakeFS) AddTo(name string, text string) error {
	fs.diary[name] += text
	return nil
}

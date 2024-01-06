package conserve

import "testing"

func TestDelete(t *testing.T) {
    executable := "conserve"
    dir := "./tmp"

    t.Cleanup(func() {
        _, err := execute("rm", "-rf", dir)
        if err != nil {
            println(err)
        }
    })
    _, err := execute(executable, "init", dir)
    if err != nil {
        t.Fatal(err)
    }

    _, err = execute(executable, "backup", dir, "./snaps")
    if err != nil {
        t.Fatal(err)
    }
    _, err = execute(executable, "backup", dir, "./snaps")
    if err != nil {
        t.Fatal(err)
    }
    _, err = execute(executable, "backup", dir, "./snaps")
    if err != nil {
        t.Fatal(err)
    }

    beforeBackups, err := Versions(executable, dir)
    if err != nil {
        t.Fatal(err)
    }
    if len(beforeBackups) != 3 {
        t.Fatalf("expected 3 backups, got %d", len(beforeBackups))
    }
    toDelete := beforeBackups[1].Name

    if err := Delete(executable, dir, toDelete); err != nil {
        t.Fatal(err)
    }

    afterBackups, err := Versions(executable, dir)
    if err != nil {
        t.Fatal(err)
    }
    if len(afterBackups) != 2 {
        t.Fatalf("expected 2 backups, got %d", len(afterBackups))
    }

    if contains(afterBackups, toDelete) {
        t.Fatalf("expected %s to be deleted", toDelete)
    }
}

func contains(collection []RawBackup, name string) bool {
    for _, item := range collection {
        if item.Name == name {
            return true
        }
    }
    return false
}

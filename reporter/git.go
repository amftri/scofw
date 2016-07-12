// TODO we need a channel for the publisher

					lastChange, alreadyTracked = lastChanges[event.Name]
	baseFolder := filepath.Join("diffs", gr.gitConfig.CurrentScoSession, filepath.Dir(event.Name))
	baseFile := filepath.Join(baseFolder, filepath.Base(event.Name))
	gr.util.CreateScoFolder(baseFolder)

	var contentA []byte
	var contentB []byte
	emptyContent := []byte("")

	options.Flags |= git.DiffShowUntrackedContent
	var contentDeltaDetermined bool

				log.Println("This file is not tracked by git", event.Name)
			contentDeltaDetermined = true
			break
		}
	}
	if !contentDeltaDetermined {
		log.Printf("ERROR: No matching git change for file: %s", event.Op, event.Name)
		log.Printf("Going to fallback - assuming this is a new file in some subdirectory")
		contentA = emptyContent
		contentB, err = ioutil.ReadFile(event.Name) // TODO can we be sure that this file is there (deleted?)?
		verifyNoError(err)
	// TODO if event.Name referes to a new file -> the patch contains "new file mode 100644" -> we should change the file mode to the original settings
	patch, err := gr.repo.PatchFromBuffers(event.Name, event.Name, contentA, contentB, &options)
	defer patch.Free()
	verifyNoError(err)
	patchString, err := patch.String()
	_, err = patch.String()
	verifyNoError(err)

	// TOOD use channel to publish change
	log.Printf("\n%s", patchString)

	// we store contentB as a snapshot of that file -> all further diffs will be made between workspace file and snapshot
	gr.util.WriteFile(&contentB, baseFile)

	gr.storeLastChange(event)
	lastChanges[event.Name] = uint32(event.Op)

	wkTree "github.com/floriangrundig/scofw/git"
	"github.com/floriangrundig/scofw/publisher"
	config                    *config.Config
	gitConfig                 *gitconfig.Config
	repo                      *git.Repository
	util                      *util.Util
	observer                  *wkTree.WorkTreeObserver
	fileEventChannel          chan *fw.FileEvent
	fileChangedMessageChannel chan *publisher.Message
func New(config *config.Config, gitConfig *gitconfig.Config, util *util.Util, observer *wkTree.WorkTreeObserver, fileEventChannel chan *fw.FileEvent, fileChangedMessageChannel chan *publisher.Message) *GitReporter {
		config:                    config,
		gitConfig:                 gitConfig,
		repo:                      repo,
		util:                      util,
		observer:                  observer,
		fileEventChannel:          fileEventChannel,
		fileChangedMessageChannel: fileChangedMessageChannel,
				break
			gr.observer.UpdateCurrentScoSession() // since we currently are not able to detect a commit we have to call update manually

					gr.handleFirstChange(event)
					gr.handleChange(event, lastChange)
				log.Println("This file is not tracked by git", event.Name)
	gr.fileChangedMessageChannel <- &publisher.Message{
		FileEvent: event,
		Patch:     &patchString,
	}
	// log.Printf("\n%s", patchString)
	// publish event
	gr.fileChangedMessageChannel <- &publisher.Message{
		FileEvent: event,
		Patch:     &patchString,
	}
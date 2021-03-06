package daemon

import (
	"github.com/neo-hu/docker_el6/distribution/reference"
	"github.com/neo-hu/docker_el6/docker/image"
)

// TagImage creates the tag specified by newTag, pointing to the image named
// imageName (alternatively, imageName can also be an image ID).
func (daemon *Daemon) TagImage(imageName, repository, tag string) (string, error) {
	imageID, _, err := daemon.GetImageIDAndOS(imageName)
	if err != nil {
		return "", err
	}

	newTag, err := reference.ParseNormalizedNamed(repository)
	if err != nil {
		return "", err
	}
	if tag != "" {
		if newTag, err = reference.WithTag(reference.TrimNamed(newTag), tag); err != nil {
			return "", err
		}
	}

	err = daemon.TagImageWithReference(imageID, newTag)
	return reference.FamiliarString(newTag), err
}

// TagImageWithReference adds the given reference to the image ID provided.
func (daemon *Daemon) TagImageWithReference(imageID image.ID, newTag reference.Named) error {
	if err := daemon.referenceStore.AddTag(newTag, imageID.Digest(), true); err != nil {
		return err
	}

	if err := daemon.imageStore.SetLastUpdated(imageID); err != nil {
		return err
	}
	daemon.LogImageEvent(imageID.String(), reference.FamiliarString(newTag), "tag")
	return nil
}

package app_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ZergsLaw/back-template/cmd/back/internal/app"
)

func TestApp_GetFile(t *testing.T) {
	t.Parallel()

	var (
		fileCache = &app.AvatarInfo{
			OwnerID: ownerID,
			FileID:  fileID,
		}
		session = app.Session{
			ID:     uuid.UUID{},
			UserID: ownerID,
		}
		file = &app.File{
			ID:          fileID,
			Name:        "name",
			ContentType: "content_type",
		}
	)

	testCases := map[string]struct {
		session             app.Session
		fileID              uuid.UUID
		repoGetFileRes      *app.AvatarInfo
		repoGetFileErr      error
		fileDownloadFileRes *app.File
		fileDownloadFileErr error
		want                *app.File
		wantErr             error
	}{
		"success":                     {session, fileID, fileCache, nil, file, nil, file, nil},
		"err_not_found_get_file":      {session, uuid.Must(uuid.NewV4()), nil, app.ErrNotFound, nil, nil, nil, app.ErrNotFound},
		"err_any_get_file":            {session, fileID, nil, errAny, nil, nil, nil, errAny},
		"err_not_found_download_file": {session, uuid.Must(uuid.NewV4()), fileCache, nil, nil, app.ErrNotFound, nil, app.ErrNotFound},
		"err_any_download_file":       {session, fileID, fileCache, nil, nil, errAny, nil, errAny},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, module, mocks, assert := start(t)

			mocks.repo.EXPECT().GetAvatar(ctx, tc.fileID).Return(tc.repoGetFileRes, tc.repoGetFileErr)

			if tc.repoGetFileErr == nil {
				mocks.file.EXPECT().DownloadFile(ctx, tc.fileID).Return(tc.fileDownloadFileRes, tc.fileDownloadFileErr)
			}

			file, err := module.GetFile(ctx, tc.session, tc.fileID)
			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.want, file)
		})
	}
}

func TestApp_RemoveAvatar(t *testing.T) {
	t.Parallel()

	var (
		fileCache1 = app.AvatarInfo{
			FileID:  fileID,
			OwnerID: ownerID,
		}
		fileCache2 = app.AvatarInfo{
			FileID:  uuid.Must(uuid.NewV4()),
			OwnerID: ownerID,
		}
		listFileCache = []app.AvatarInfo{fileCache2}
		user1         = app.User{
			ID:       ownerID,
			Email:    "test@test.com",
			Name:     "name",
			AvatarID: uuid.Must(uuid.NewV4()),
		}
		session = app.Session{
			ID:     uuid.Must(uuid.NewV4()),
			UserID: ownerID,
		}
		sessionAnother = app.Session{
			ID:     uuid.Must(uuid.NewV4()),
			UserID: uuid.Must(uuid.NewV4()),
		}
		user2 = user1
	)

	testCases := map[string]struct {
		session                        app.Session
		fileID                         uuid.UUID
		repoGetFileRes                 *app.AvatarInfo
		repoGetFileErr                 error
		repoDeleteAvatarCacheErr       error
		fileDeleteFileErr              error
		repoListAvatarCacheByUserIDRes []app.AvatarInfo
		repoListAvatarCacheByUserIDErr error
		repoByIDRes                    *app.User
		repoByIDErr                    error
		repoUpdateRes                  *app.User
		repoUpdateErr                  error
		want                           error
	}{
		"success":                                   {session, fileID, &fileCache1, nil, nil, nil, listFileCache, nil, &user1, nil, &app.User{}, nil, nil},
		"err_access_denied":                         {sessionAnother, fileID, &fileCache1, nil, nil, nil, nil, nil, nil, nil, &app.User{}, nil, app.ErrAccessDenied},
		"err_any_repo_get_file":                     {session, fileID, &fileCache1, errAny, nil, nil, nil, nil, nil, nil, &app.User{}, nil, errAny},
		"err_any_repo_delete_avatar_cache":          {session, fileID, &fileCache1, nil, errAny, nil, nil, nil, nil, nil, &app.User{}, nil, errAny},
		"err_any_file_delete_file":                  {session, fileID, &fileCache1, nil, nil, errAny, nil, nil, nil, nil, &app.User{}, nil, errAny},
		"err_any_repo_list_avatar_cache_by_user_id": {session, fileID, &fileCache1, nil, nil, nil, nil, errAny, nil, nil, &app.User{}, nil, errAny},
		"err_any_repo_by_id":                        {session, fileID, &fileCache1, nil, nil, nil, listFileCache, nil, nil, errAny, &app.User{}, nil, errAny},
		"err_any_repo_update":                       {session, fileID, &fileCache1, nil, nil, nil, listFileCache, nil, &user2, nil, &app.User{}, errAny, errAny},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, module, mocks, assert := start(t)

			mocks.repo.EXPECT().GetAvatar(ctx, tc.fileID).Return(tc.repoGetFileRes, tc.repoGetFileErr)

			if tc.repoGetFileErr == nil && !errors.Is(tc.want, app.ErrAccessDenied) {
				mocks.repo.EXPECT().Tx(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(repo app.Repo) error) error {
					return fn(mocks.repo)
				})

				mocks.repo.EXPECT().DeleteAvatar(ctx, tc.session.UserID, tc.fileID).Return(tc.repoDeleteAvatarCacheErr)

				if tc.repoDeleteAvatarCacheErr == nil {
					mocks.file.EXPECT().DeleteFile(ctx, tc.fileID).Return(tc.fileDeleteFileErr)
				}

				if tc.repoDeleteAvatarCacheErr == nil && tc.fileDeleteFileErr == nil {
					mocks.repo.EXPECT().ListAvatarByUserID(ctx, tc.session.UserID).Return(tc.repoListAvatarCacheByUserIDRes, tc.repoListAvatarCacheByUserIDErr)
				}

				if tc.repoDeleteAvatarCacheErr == nil && tc.fileDeleteFileErr == nil && tc.repoListAvatarCacheByUserIDErr == nil {
					mocks.repo.EXPECT().UserByID(ctx, tc.session.UserID).Return(tc.repoByIDRes, tc.repoByIDErr)
				}

				if tc.repoDeleteAvatarCacheErr == nil && tc.fileDeleteFileErr == nil &&
					tc.repoListAvatarCacheByUserIDErr == nil && tc.repoByIDErr == nil {
					newAvatarID := uuid.Nil
					if len(tc.repoListAvatarCacheByUserIDRes) > 0 {
						newAvatarID = tc.repoListAvatarCacheByUserIDRes[0].FileID
					}
					tc.repoByIDRes.AvatarID = newAvatarID
					mocks.repo.EXPECT().UserUpdate(ctx, *tc.repoByIDRes).Return(tc.repoUpdateRes, tc.repoUpdateErr)
				}
			}

			err := module.RemoveAvatar(ctx, tc.session, tc.fileID)
			assert.ErrorIs(err, tc.want)
		})
	}
}

func TestApp_AddAvatar(t *testing.T) {
	t.Parallel()

	file, err := os.Open(pngFilePath)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, file.Close()) })
	fileInfo, err := file.Stat()
	require.NoError(t, err)

	var (
		f = &app.File{
			ID:             fileID,
			UserID:         ownerID,
			Name:           fileInfo.Name(),
			ContentType:    "image/jpeg",
			Size:           fileInfo.Size(),
			ReadSeekCloser: file,
		}
		f2 = &app.File{
			ID:             fileID,
			UserID:         ownerID,
			Name:           fileInfo.Name(),
			ContentType:    "image",
			Size:           fileInfo.Size(),
			ReadSeekCloser: file,
		}
		fileCache1 = app.AvatarInfo{
			FileID:  fileID,
			OwnerID: ownerID,
		}
		user1 = app.User{
			ID:       ownerID,
			Email:    "test@test.com",
			Name:     "name",
			AvatarID: fileID,
		}
		session = app.Session{
			ID:     uuid.Must(uuid.NewV4()),
			UserID: ownerID,
		}
		sessionAnother = app.Session{
			ID:     uuid.Must(uuid.NewV4()),
			UserID: uuid.Must(uuid.NewV4()),
		}
	)
	testCases := map[string]struct {
		session         app.Session
		fileID          uuid.UUID
		avatarInfo      app.AvatarInfo
		accessDenied    error
		downloadFileRes *app.File
		downloadFileErr error
		getCountRes     int
		getCountErr     error
		getUserRes      *app.User
		getUserErr      error
		saveAvatarErr   error
		userUpdateErr   error
		unknownFormat   error
		maxAvatarErr    error
		wantErr         error
	}{
		"success":                {session: session, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f, downloadFileErr: nil, getCountRes: 0, getCountErr: nil, getUserRes: lo.ToPtr(user1), saveAvatarErr: nil, userUpdateErr: nil, getUserErr: nil, unknownFormat: nil, maxAvatarErr: nil, wantErr: nil},
		"a.file.DownloadFile":    {session: session, fileID: fileID, avatarInfo: app.AvatarInfo{}, downloadFileRes: nil, downloadFileErr: errAny, getCountRes: 0, getCountErr: nil, getUserRes: lo.ToPtr(user1), saveAvatarErr: nil, userUpdateErr: nil, getUserErr: nil, unknownFormat: nil, maxAvatarErr: nil, wantErr: errAny},
		"validateFormat":         {session: session, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f2, downloadFileErr: nil, getCountRes: 0, getCountErr: nil, getUserRes: nil, getUserErr: nil, saveAvatarErr: nil, userUpdateErr: nil, unknownFormat: app.ErrInvalidImageFormat, maxAvatarErr: nil, wantErr: app.ErrInvalidImageFormat},
		"a.repo.GetCountAvatars": {session: session, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f, downloadFileErr: nil, getCountRes: 0, getCountErr: errAny, getUserRes: nil, getUserErr: nil, saveAvatarErr: nil, userUpdateErr: nil, unknownFormat: nil, maxAvatarErr: nil, wantErr: errAny},
		"avatars_limit":          {session: session, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f, downloadFileErr: nil, getCountRes: 30, getCountErr: nil, getUserRes: nil, getUserErr: nil, saveAvatarErr: nil, userUpdateErr: nil, unknownFormat: nil, maxAvatarErr: app.ErrMaxFiles, wantErr: app.ErrMaxFiles},
		"repo.UserByID":          {session: session, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f, downloadFileErr: nil, getCountRes: 0, getCountErr: nil, getUserRes: lo.ToPtr(user1), saveAvatarErr: nil, userUpdateErr: nil, getUserErr: nil, unknownFormat: nil, maxAvatarErr: nil, wantErr: nil},
		"a.repo.SaveAvatar":      {session: session, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f, downloadFileErr: nil, getCountRes: 0, getCountErr: nil, getUserRes: lo.ToPtr(user1), saveAvatarErr: errAny, userUpdateErr: nil, getUserErr: nil, unknownFormat: nil, maxAvatarErr: nil, wantErr: errAny},
		"repo.UserUpdate":        {session: session, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f, downloadFileErr: nil, getCountRes: 0, getCountErr: nil, getUserRes: lo.ToPtr(user1), saveAvatarErr: nil, userUpdateErr: errAny, getUserErr: nil, unknownFormat: nil, maxAvatarErr: nil, wantErr: errAny},
		"accessDenied":           {session: sessionAnother, fileID: fileID, avatarInfo: fileCache1, downloadFileRes: f, downloadFileErr: nil, getCountRes: 0, getCountErr: nil, getUserRes: lo.ToPtr(user1), saveAvatarErr: nil, userUpdateErr: errAny, getUserErr: nil, unknownFormat: nil, maxAvatarErr: nil, wantErr: app.ErrAccessDenied},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, module, mocks, assert := start(t)

			mocks.file.EXPECT().DownloadFile(gomock.Any(), tc.fileID).Return(tc.downloadFileRes, tc.downloadFileErr)

			var splits []string
			if tc.downloadFileErr == nil {
				splits = strings.Split(tc.downloadFileRes.ContentType, "/")
			}
			if len(splits) >= 2 && !errors.Is(tc.wantErr, app.ErrAccessDenied) {
				if splits[1] == "jpeg" {
					mocks.repo.EXPECT().Tx(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(repo app.Repo) error) error {
						return fn(mocks.repo)
					})

					mocks.repo.EXPECT().GetCountAvatars(ctx, tc.session.UserID).Return(tc.getCountRes, tc.getCountErr)

					if (tc.getCountErr == nil || errors.Is(tc.getCountErr, app.ErrNotFound)) && tc.getCountRes < 10 {
						mocks.repo.EXPECT().UserByID(ctx, tc.session.UserID).Return(tc.getUserRes, tc.getUserErr)
					}

					if (tc.getCountErr == nil || errors.Is(tc.getCountErr, app.ErrNotFound)) && tc.getCountRes < 10 && tc.getUserErr == nil {
						mocks.repo.EXPECT().SaveAvatar(ctx, tc.avatarInfo).Return(tc.saveAvatarErr)
					}

					if (tc.getCountErr == nil || errors.Is(tc.getCountErr, app.ErrNotFound)) && tc.getCountRes < 10 && tc.getUserErr == nil && tc.saveAvatarErr == nil {
						mocks.repo.EXPECT().UserUpdate(ctx, *tc.getUserRes).Return(tc.getUserRes, tc.userUpdateErr)
					}

				}
			}

			err := module.AddAvatar(ctx, tc.session, tc.fileID)
			assert.ErrorIs(err, tc.wantErr)
		})
	}
}
